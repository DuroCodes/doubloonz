package main

type project struct {
	Name      string  `json:"name"`
	Doubloons uint64  `json:"doubloons"`
	Hours     float64 `json:"hours"`
}

type storage struct {
	Region    string    `json:"region"`
	Projects  []project `json:"projects"`
}

type prize struct {
	Name string `json:"name"`
	Cost uint   `json:"cost"`
}

type mode string

const (
	modeAddProject mode = "Add Project"
	modeDelete     mode = "Delete Project"
	modeEdit       mode = "Edit Project"
	modePrize      mode = "Select Prize"
	modeRegion     mode = "Change Region"
	modeExit       mode = "Exit"
)
