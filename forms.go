package main

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
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
						return errors.New("hours must be a number")
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

func prizeSelection() (prize, error) {
	var selectedPrize prize

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[prize]().
				Title("Choose your prize").
				Options(
					huh.NewOption("Prize 1", prize{"Prize 1", 100}),
					huh.NewOption("Prize 2", prize{"Prize 2", 200}),
				).
				Value(&selectedPrize),
		),
	)

	err := form.Run()
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
