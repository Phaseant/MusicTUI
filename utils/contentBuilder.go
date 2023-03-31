package utils

import (
	"fmt"
	"strings"

	"github.com/Phaseant/MusicTUI"
	"github.com/charmbracelet/lipgloss"
)

// style definitions
var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2).
		Height(8).
		Width(40)

	listItem = lipgloss.NewStyle().PaddingLeft(2)

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(false).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss").
			Align(lipgloss.Left).
			Bold(true)

	authorStyle = lipgloss.NewStyle().Align(lipgloss.Left)

	descStyle = lipgloss.NewStyle().MarginTop(1).Italic(true)
)

func ContentBuilder(album MusicTUI.Album) string {
	doc := strings.Builder{}

	//title
	{
		title := lipgloss.JoinVertical(
			lipgloss.Top,
			titleStyle.Render(album.Title),
			authorStyle.Render(album.Author),
		)

		doc.WriteString(title)
	}

	{
		desc := descStyle.Render(album.Description)

		doc.WriteString(desc)
	}

	{
		list := list.Render(listItems(album.Songs)...)

		doc.WriteString(list)
	}

	return doc.String()

}

func listItems(songs []MusicTUI.Song) []string {
	var result []string
	curStr := strings.Builder{}
	for i, el := range songs {
		song := listItem.Render(el.Title)
		curStr.WriteString(fmt.Sprint(i) + ". " + song)
		result = append(result, curStr.String())
	}

	return result
}
