package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// safe exit
		<-sigChan
		os.Exit(0)
	}()

	projects, err := loadProjects()
	if err != nil {
		log.Fatal(err)
	}

	var totalDoubloons uint64
	var totalHours float64

	for _, p := range projects {
		totalDoubloons += p.Doubloons
		totalHours += p.Hours
	}

	for {
		fmt.Print("\033[H\033[2J")

		if len(projects) == 0 {
			displayProjects([]project{
				{
					Name:      "No projects yet!",
					Doubloons: 0,
					Hours:     0,
				},
			})
		}

		if len(projects) > 0 {
			displayProjects(projects)
		}

		selectedMode, err := selectMode()
		if err != nil {
			return
		}

		switch selectedMode {
		case modeAddProject:
			project, err := newProject()
			if err != nil {
				continue
			}
			projects = append(projects, project)
			totalDoubloons += project.Doubloons
			totalHours += project.Hours
			saveProjects(projects)

		case modeDelete:
			if newProjects, ok := deleteProject(projects); ok {
				totalDoubloons = 0
				totalHours = 0
				for _, p := range newProjects {
					totalDoubloons += p.Doubloons
					totalHours += p.Hours
				}
				projects = newProjects
				saveProjects(projects)
			}

		case modePrize:
			if len(projects) == 0 {
				fmt.Println("Add some projects first!")
				continue
			}
			selectedPrize, err := prizeSelection()
			if err != nil {
				continue
			}
			averageHourlyRate := float64(totalDoubloons) / totalHours

			messageStyle := lipgloss.NewStyle().
				Padding(0, 1).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("57"))

			if totalDoubloons >= uint64(selectedPrize.cost) {
				msg := fmt.Sprintf("You can afford %s!", selectedPrize.name)
				fmt.Println(messageStyle.Render(msg))
			} else {
				doubloonsNeeded := uint64(selectedPrize.cost) - totalDoubloons
				hoursNeeded := float64(doubloonsNeeded) / averageHourlyRate
				msg := fmt.Sprintf("To afford %s, you need:\n"+
					"•   %d more doubloons\n"+
					"•   ≈%.1f more hours ",
					selectedPrize.name, doubloonsNeeded, hoursNeeded)
				fmt.Println(messageStyle.Render(msg))
			}
			fmt.Println("\nPress Enter to continue...")
			fmt.Scanln()

		case modeEdit:
			if newProjects, ok := editProject(projects); ok {
				totalDoubloons = 0
				totalHours = 0
				for _, p := range newProjects {
					totalDoubloons += p.Doubloons
					totalHours += p.Hours
				}

				projects = newProjects
				saveProjects(projects)
			}

		case modeExit:
			os.Exit(0)
		}
	}
}
