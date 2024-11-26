package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func displayProjects(state storage) {
	columns := []table.Column{
		{Title: "Project Name", Width: 29}, // odd so the subtext can be cenetered
		{Title: "Doubloons", Width: 15},
		{Title: "Hours", Width: 12},
		{Title: "Rate", Width: 20},
	}

	rows := []table.Row{}
	for _, p := range state.Projects {
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

	totalDoubloons := uint64(0)
	totalHours := float64(0)
	for _, p := range state.Projects {
		totalDoubloons += p.Doubloons
		totalHours += p.Hours
	}
	avgRate := fmt.Sprintf("%.2f", float64(totalDoubloons)/totalHours)
	if totalHours == 0 {
		avgRate = "N/A"
	}

	tableWidth := lipgloss.Width(t.View())

	subtextStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Align(lipgloss.Center).
		Width(tableWidth)

	subtext := fmt.Sprintf("%d   •   %.2f hrs • ≈%s  /hr •   %s\n", totalDoubloons, totalHours, avgRate, state.Region)

	fmt.Println(baseStyle.Render(t.View()))
	fmt.Println(subtextStyle.Render(subtext))
}
