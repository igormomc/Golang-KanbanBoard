package main

import "github.com/charmbracelet/lipgloss"

var (
	columnStyle  = lipgloss.NewStyle().Padding(1, 2).Foreground(lipgloss.Color("#ADD8E6")) // Light Blue text color
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#87CEEB")). // Sky Blue border color
			Foreground(lipgloss.Color("#ADD8E6"))        // Light Blue text color

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
)
