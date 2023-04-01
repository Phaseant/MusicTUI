package view

import (
	"github.com/Phaseant/MusicTUI"
	"github.com/Phaseant/MusicTUI/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

type ListModel struct {
	list list.Model
}

type item struct {
	Album       MusicTUI.Album
	title       string
	description string
}

func InitDelegateStyle() list.DefaultItemStyles {
	style := list.DefaultItemStyles{}

	style.SelectedTitle.Foreground(YellowColor)
	style.SelectedDesc.Foreground(YellowColor)

	style.FilterMatch.Foreground(YellowColor)

	return style
}

var listStyle = lipgloss.NewStyle().Background(YellowColor).Foreground(BlackColor)

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

func InitListView() (tea.Model, tea.Cmd) {

	//TODO: if is not able to find connect to server, show downloading spinner
	albums, err := utils.FetchAlbums("http://localhost:8000/api/album")
	if err != nil {
		log.Errorf("unable to fetch data: %v", err)
	}

	var items []list.Item

	for _, el := range albums {
		items = append(items, item{Album: el, title: el.Title, description: el.Author})
	}

	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle.Foreground(YellowColor)
	delegate.Styles.SelectedDesc.Foreground(YellowColor)

	delegate.Styles.FilterMatch.Foreground(YellowColor)

	// setup list
	m := ListModel{list: list.New(items, delegate, 10, 10)}

	m.list.Styles.Title = listStyle
	w, h := docStyle.GetFrameSize()
	if WindowSize.Width != 0 {
		m.list.SetSize(WindowSize.Width-w, WindowSize.Height-h)
	}
	m.list.Title = "Albums"
	return m, nil
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w, h := docStyle.GetFrameSize()
		WindowSize = msg
		m.list.SetSize(msg.Width-w, msg.Height-h)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case msg.String() == "enter":
			curItem := m.list.SelectedItem().(item)
			albumView := InitAlbumView(curItem.Album)
			return albumView.Update(WindowSize)
		}

	}

	m.list, cmd = m.list.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m ListModel) View() string {
	return docStyle.Render(m.list.View() + "\n")
}
