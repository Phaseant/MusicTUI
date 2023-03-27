package main

import (
	"github.com/Phaseant/MusicTUI/view"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

func main() {
	m, _ := view.InitListView()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Panicf("unable to start tea program: %v", err)
	}
}
