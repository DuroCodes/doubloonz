package main

import (
	"encoding/json"
	"os"
)

func saveProjects(projects []project) error {
	data, err := json.Marshal(projects)
	if err != nil {
		return err
	}
	return os.WriteFile("projects.json", data, 0644)
}

func loadProjects() ([]project, error) {
	data, err := os.ReadFile("projects.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []project{}, nil
		}
		return nil, err
	}
	var projects []project
	err = json.Unmarshal(data, &projects)
	return projects, err
}

func loadPrizes() ([]prize, error) {
	data, err := os.ReadFile("prizes.json")
	if err != nil {
		return nil, err
	}
	var prizes []prize
	err = json.Unmarshal(data, &prizes)
	return prizes, err
}
