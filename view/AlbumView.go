package view

import (
	"github.com/Phaseant/MusicTUI"
	"github.com/Phaseant/MusicTUI/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AlbumModel struct {
	viewport viewport.Model
	album    MusicTUI.Album
	content  string
}

func InitAlbumView(album MusicTUI.Album) *AlbumModel {
	m := AlbumModel{album: album, content: utils.ContentBuilder(album)}

	m.viewport = viewport.New(20, 20)

	// if WindowSize.Width != 0 {
	// 	m.viewport.Height = WindowSize.Height
	// 	m.viewport.Width = WindowSize.Width
	// }

	m.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)
	m.viewport.SetContent(m.content)
	return &m
}

func (m AlbumModel) Init() tea.Cmd {
	return nil
}

func (m AlbumModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" {
			return m, tea.Quit
		}

		if k := msg.String(); k == "esc" {
			return InitListView()
		}

	case tea.WindowSizeMsg:
		h, w := docStyle.GetFrameSize()
		m.viewport.Height = msg.Height - h
		m.viewport.Width = msg.Width - w
		WindowSize = msg
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m AlbumModel) View() string {
	return m.viewport.View()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
