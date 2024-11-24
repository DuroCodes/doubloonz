package main

type project struct {
	Name      string  `json:"name"`
	Doubloons uint64  `json:"doubloons"`
	Hours     float64 `json:"hours"`
}

type prize struct {
	name string
	cost uint
}

type mode string

const (
	modeAddProject mode = "Add Project"
	modeDelete     mode = "Delete Project"
	modeEdit       mode = "Edit Project"
	modePrize      mode = "Select Prize"
	modeExit       mode = "Exit"
)