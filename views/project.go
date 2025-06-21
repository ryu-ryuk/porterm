package views

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"porterm/internal/data"
	"porterm/styles"
)

func Projects(width int) string {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(data.GlamourStyle),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return styles.ErrorStyle.Render("Renderer error: " + err.Error())
	}

	rendered, err := renderer.Render(data.ProjectsData)
	if err != nil {
		return styles.ErrorStyle.Render("Markdown error: " + err.Error())
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.Heading.Render("Projects"),
		rendered,
		"",
		styles.Help.Render("Press esc to return to menu."),
	)
}
