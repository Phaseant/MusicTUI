package utils

import (
	"fmt"
	"strings"

	"github.com/Phaseant/MusicTUI"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// style definitions
var (
	YellowColor = lipgloss.Color("#FCCB06")
	BlackColor  = lipgloss.Color("#000000")

	url = lipgloss.NewStyle().Render

	docStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, true).Margin(1)

	listStyle = lipgloss.NewStyle()

	listItem = lipgloss.NewStyle().PaddingLeft(2)

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#3EB489")).
			Align(lipgloss.Top).
			Bold(true).
			Foreground(YellowColor)

	authorStyle = lipgloss.NewStyle().Align(lipgloss.Top).MarginLeft(1).Padding(0, 1)

	descStyle = lipgloss.NewStyle().MarginTop(1).Italic(true).Border(lipgloss.NormalBorder()).BorderTop(true)

	HeaderStyle = lipgloss.NewStyle().Background(YellowColor).MarginLeft(1).Foreground(BlackColor).MarginBottom(1)

	descHeaderStyle = lipgloss.NewStyle().Background(YellowColor).Foreground(BlackColor).MarginBottom(1)

	GeniusHeaderStyle = lipgloss.NewStyle().Background(YellowColor).MarginLeft(1).Foreground(BlackColor)
)

func ContentBuilder(album MusicTUI.Album, msg tea.Msg) string {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		docStyle.Width(msg.Width - 7)
		docStyle.Height(msg.Height - 7)

		descStyle.Width(msg.Width - 9)
	}

	doc := strings.Builder{}

	titleAu := lipgloss.JoinHorizontal(lipgloss.Center, titleStyle.Render(album.Title), authorStyle.Render(album.Author), authorStyle.Render(fmt.Sprint(album.Year)))

	doc.WriteString(titleAu)
	doc.WriteString("\n")
	doc.WriteString("\n")

	doc.WriteString(GeniusHeaderStyle.Render("Genius:") + " " + url(album.GeniousLink))

	doc.WriteString("\n")

	// doc.WriteString(HeaderStyle.Render("Description"))
	desc := descStyle.Render(descHeaderStyle.Render("Description") + "\n" + album.Description)

	doc.WriteString(desc)
	doc.WriteString("\n")
	doc.WriteString("\n")

	list := listStyle.Render(listItems(album.Songs))
	doc.WriteString(HeaderStyle.Render("Songs"))
	doc.WriteString("\n")
	doc.WriteString(list)
	doc.WriteString("\n\n")

	doc.WriteString(HeaderStyle.Render("Album duration: " + album.Duration))

	return docStyle.Render(doc.String())

}

func listItems(songs []MusicTUI.Song) string {
	curStr := strings.Builder{}
	size := len(songs)
	for i, el := range songs {
		song := listItem.Render(fmt.Sprint(i+1) + ". " + el.Title + " " + el.Duration)
		curStr.WriteString(song)
		if (i + 1) != size {
			curStr.WriteString("\n\n")
		}
	}

	return curStr.String()
}
