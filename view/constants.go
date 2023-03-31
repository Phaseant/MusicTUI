package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	WindowSize tea.WindowSizeMsg
	docStyle   lipgloss.Style = lipgloss.NewStyle().Margin(2)
)
