package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func displayProjects(projects []project) {
	columns := []table.Column{
		{Title: "Project Name", Width: 30},
		{Title: "Doubloons", Width: 15},
		{Title: "Hours", Width: 12},
		{Title: "Rate", Width: 20},
	}

	rows := []table.Row{}
	for _, p := range projects {
		var rateStr string
		if p.Hours == 0 {
			rateStr = "N/A  /hr"
		} else {
			rate := float64(p.Doubloons) / p.Hours
			rateStr = fmt.Sprintf("%.2f  /hr", rate)
		}
		rows = append(rows, table.Row{
			p.Name,
			fmt.Sprintf("%d ", p.Doubloons),
			fmt.Sprintf("%.2f", p.Hours),
			rateStr,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(rows)+1),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		UnsetBackground().
		UnsetForeground().
		Bold(false)
	t.SetStyles(s)

	baseStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	fmt.Println(baseStyle.Render(t.View()) + "\n")
}
