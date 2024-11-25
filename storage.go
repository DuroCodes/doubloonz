package main

import (
	"encoding/json"
	"os"
)

func saveProjects(store storage) error {
	data, err := json.Marshal(store)
	if err != nil {
		return err
	}
	return os.WriteFile("projects.json", data, 0644)
}

func loadProjects() (storage, error) {
	data, err := os.ReadFile("projects.json")
	if err != nil {
		if os.IsNotExist(err) {
			return storage{Region: "US", Projects: []project{}}, nil
		}
		return storage{}, err
	}

	var store storage
	err = json.Unmarshal(data, &store)
	if err == nil && store.Region != "" {
		return store, nil
	}

	// backwards compat
	var oldProjects []project
	err = json.Unmarshal(data, &oldProjects)
	if err == nil {
		return storage{
			Region:   "US",
			Projects: oldProjects,
		}, nil
	}

	return storage{}, err
}

func loadPrizes() (map[string][]prize, error) {
	data, err := os.ReadFile("prizes.json")
	if err != nil {
		return nil, err
	}
	var prizes map[string][]prize
	err = json.Unmarshal(data, &prizes)
	return prizes, err
}
