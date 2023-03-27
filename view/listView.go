package view

import (
	"github.com/Phaseant/MusicTUI"
	"github.com/Phaseant/MusicTUI/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	list    list.Model
	curItem int
}

type item struct {
	Album       MusicTUI.Album
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

func InitListView() (tea.Model, tea.Msg) {
	albums, err := utils.FetchAlbums("http://localhost:8000/api/album")
	if err != nil {
		log.Errorf("unable to fetch data: %v", err)
	}

	var items []list.Item

	for _, el := range albums {
		items = append(items, item{Album: el, title: el.Title, description: el.Author})
	}

	h, w := docStyle.GetFrameSize()

	m := Model{list: list.New(items, list.NewDefaultDelegate(), w, h)}

	m.list.Title = "Albums"

	return m, func() tea.Msg { return nil }
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		// case "enter":

		// }
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return docStyle.Render(m.list.View() + "\n")
}
