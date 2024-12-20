package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func newProject() (project, error) {
	var name, doubloons, hours string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Value(&name).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("project name cannot be empty")
					}
					return nil
				}),

			huh.NewInput().
				Title("Doubloons Earned").
				Value(&doubloons).
				Validate(func(str string) error {
					_, err := strconv.ParseUint(str, 10, 32)
					if err != nil {
						return errors.New("doubloons must be a number")
					}
					return nil
				}),

			huh.NewInput().
				Title("Hours Spent").
				Value(&hours).
				Validate(func(str string) error {
					_, err := strconv.ParseFloat(str, 32)
					if err != nil {
						return errors.New("hours must be a number")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		return project{}, err
	}

	doubloons_uint, _ := strconv.ParseUint(doubloons, 10, 32)
	hours_float, _ := strconv.ParseFloat(hours, 32)

	return project{
		Name:      name,
		Doubloons: doubloons_uint,
		Hours:     hours_float,
	}, nil
}

func prizeSelection(region string) (prize, error) {
	var selectedPrize prize

	loadedPrizes, err := loadPrizes()
	if err != nil {
		return prize{}, err
	}

	prizes := make([]huh.Option[prize], len(loadedPrizes[region]))
	for i, p := range loadedPrizes[region] {
		prizes[i] = huh.NewOption(p.Name, p)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[prize]().
				Title("Choose your prize").
				Options(prizes...).
				Height(10).
				Value(&selectedPrize),
		),
	)

	err = form.Run()
	if err != nil {
		return prize{}, err
	}

	return selectedPrize, nil
}

func deleteProject(projects []project) ([]project, bool) {
	if len(projects) == 0 {
		return projects, false
	}

	var selectedIndex int
	options := make([]huh.Option[int], len(projects))
	for i, p := range projects {
		options[i] = huh.NewOption(p.Name, i)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose project to delete").
				Options(options...).
				Value(&selectedIndex),
		),
	)

	err := form.Run()
	if err != nil {
		return projects, false
	}

	return append(projects[:selectedIndex], projects[selectedIndex+1:]...), true
}

func editProject(projects []project) ([]project, bool) {
	if len(projects) == 0 {
		return projects, false
	}

	var selectedIndex int
	options := make([]huh.Option[int], len(projects))
	for i, p := range projects {
		options[i] = huh.NewOption(p.Name, i)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose project to edit").
				Options(options...).
				Value(&selectedIndex),
		),
	)

	err := form.Run()
	if err != nil {
		return projects, false
	}

	var name, doubloons, hours string
	editForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Value(&name).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("project name cannot be empty")
					}
					return nil
				}),

			huh.NewInput().
				Title("Doubloons Earned").
				Value(&doubloons).
				Validate(func(str string) error {
					_, err := strconv.ParseUint(str, 10, 32)
					if err != nil {
						return errors.New("doubloons must be a number")
					}
					return nil
				}),

			huh.NewInput().
				Title("Hours Spent").
				Value(&hours).
				Validate(func(str string) error {
					_, err := strconv.ParseFloat(str, 32)
					if err != nil {
						return errors.New("invalid hours")
					}
					return nil
				}),
		),
	)

	if err := editForm.Run(); err != nil {
		return projects, false
	}

	// safe to ignore errors since input already validated
	doubloons_uint, _ := strconv.ParseUint(doubloons, 10, 32)
	hours_float, _ := strconv.ParseFloat(hours, 32)

	projects[selectedIndex] = project{
		Name:      name,
		Doubloons: doubloons_uint,
		Hours:     hours_float,
	}

	return projects, true
}

func selectMode() (mode, error) {
	var selectedMode mode

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[mode]().
				Title("Choose mode").
				Options(
					huh.NewOption("Add Project", modeAddProject),
					huh.NewOption("Edit Project", modeEdit),
					huh.NewOption("Delete Project", modeDelete),
					huh.NewOption("Select Prize", modePrize),
					huh.NewOption("Calculate", modeCalc),
					huh.NewOption("Change Region", modeRegion),
					huh.NewOption("Exit", modeExit),
				).
				Filtering(false).
				Value(&selectedMode),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return selectedMode, nil
}

func selectRegion(currentRegion string) (string, error) {
	var selectedRegion string = currentRegion

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose region").
				Options(
					huh.NewOption("United States", "US"),
					huh.NewOption("European Union", "EU"),
					huh.NewOption("India", "IN"),
					huh.NewOption("Canada", "CA"),
					huh.NewOption("Worldwide", "WW"),
				).
				Value(&selectedRegion),
		),
	)

	err := form.Run()
	if err != nil {
		return currentRegion, err
	}

	return selectedRegion, nil
}

func calculateMode(totalDoubloons uint64, totalHours float64) {
	avgRate := float64(totalDoubloons) / totalHours
	if totalHours == 0 {
		avgRate = 0
	}

	messageStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("57"))

	var calcMode string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to calculate?").
				Options(
					huh.NewOption("Hours needed for doubloons", "hours"),
					huh.NewOption("Doubloons from hours", "doubloons"),
				).
				Value(&calcMode),
		),
	)

	if err := form.Run(); err != nil {
		return
	}

	switch calcMode {
	case "hours":
		var targetDoubloons string
		hoursForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("How many doubloons do you want?").
					Validate(func(str string) error {
						_, err := strconv.ParseUint(str, 10, 64)
						if err != nil {
							return errors.New("doubloons must be a number")
						}
						return nil
					}).
					Value(&targetDoubloons),
			),
		)

		if err := hoursForm.Run(); err != nil {
			return
		}

		target, _ := strconv.ParseUint(targetDoubloons, 10, 64)
		hoursNeeded := float64(target) / avgRate

		msg := fmt.Sprintf("%d  / ≈%.2f  /hr = %.1f hours", target, avgRate, hoursNeeded)
		fmt.Println(messageStyle.Render(msg))

	case "doubloons":
		var targetHours string
		doubloonForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("How many hours will you work?").
					Validate(func(str string) error {
						_, err := strconv.ParseFloat(str, 64)
						if err != nil {
							return errors.New("invalid hours")
						}
						return nil
					}).
					Value(&targetHours),
			),
		)

		if err := doubloonForm.Run(); err != nil {
			return
		}

		hours, _ := strconv.ParseFloat(targetHours, 64)
		msg := fmt.Sprintf("%.2f  /hr × %.1f hours = %.0f ", avgRate, hours, hours*avgRate)

		fmt.Println(messageStyle.Render(msg))
	}

	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
}
