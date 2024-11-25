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

	state, err := loadProjects()
	if err != nil {
		log.Fatal(err)
	}

	var totalDoubloons uint64
	var totalHours float64

	for _, p := range state.Projects {
		totalDoubloons += p.Doubloons
		totalHours += p.Hours
	}

	for {
		fmt.Print("\033[H\033[2J")

		if len(state.Projects) == 0 {
			displayProjects([]project{
				{
					Name:      "No projects yet!",
					Doubloons: 0,
					Hours:     0,
				},
			})
		}

		if len(state.Projects) > 0 {
			displayProjects(state.Projects)
		}

		fmt.Printf("Current region: %s\n\n", state.Region)

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
			state.Projects = append(state.Projects, project)
			totalDoubloons += project.Doubloons
			totalHours += project.Hours
			saveProjects(state)

		case modeDelete:
			if newProjects, ok := deleteProject(state.Projects); ok {
				totalDoubloons = 0
				totalHours = 0
				for _, p := range newProjects {
					totalDoubloons += p.Doubloons
					totalHours += p.Hours
				}
				state.Projects = newProjects
				saveProjects(state)
			}

		case modePrize:
			if len(state.Projects) == 0 {
				fmt.Println("Add some projects first!")
				continue
			}
			selectedPrize, err := prizeSelection(state.Region)
			if err != nil {
				continue
			}
			averageHourlyRate := float64(totalDoubloons) / totalHours

			messageStyle := lipgloss.NewStyle().
				Padding(0, 1).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("57"))

			if totalDoubloons >= uint64(selectedPrize.Cost) {
				msg := fmt.Sprintf("You can afford %s!", selectedPrize.Name)
				fmt.Println(messageStyle.Render(msg))
			} else {
				doubloonsNeeded := uint64(selectedPrize.Cost) - totalDoubloons
				hoursNeeded := float64(doubloonsNeeded) / averageHourlyRate
				msg := fmt.Sprintf("To afford %s, you need:\n"+
					"•   %d more doubloons\n"+
					"•   ≈%.1f more hours ",
					selectedPrize.Name, doubloonsNeeded, hoursNeeded)
				fmt.Println(messageStyle.Render(msg))
			}
			fmt.Println("\nPress Enter to continue...")
			fmt.Scanln()

		case modeEdit:
			if newProjects, ok := editProject(state.Projects); ok {
				totalDoubloons = 0
				totalHours = 0
				for _, p := range newProjects {
					totalDoubloons += p.Doubloons
					totalHours += p.Hours
				}

				state.Projects = newProjects
				saveProjects(state)
			}

		case modeRegion:
			newRegion, err := selectRegion(state.Region)
			if err != nil {
				continue
			}
			if newRegion != state.Region {
				state.Region = newRegion
				saveProjects(state)
			}

		case modeExit:
			os.Exit(0)
		}
	}
}
