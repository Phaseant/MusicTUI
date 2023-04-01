package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	WindowSize  tea.WindowSizeMsg
	docStyle    lipgloss.Style = lipgloss.NewStyle().Margin(2)
	YellowColor                = lipgloss.Color("#FCCB06")
	BlackColor                 = lipgloss.Color("#000000")
)
