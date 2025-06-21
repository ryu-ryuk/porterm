package views

import (
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"porterm/styles"
)

func Projects(width int) string {
	// read projects from file
	data, err := os.ReadFile("assets/paste.txt")
	if err != nil {
		return styles.ErrorStyle.Render("Error reading projects: " + err.Error())
	}

	// render markdown with Glamour
	renderer, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(width),
		glamour.WithStylesFromJSONFile("assets/glamour-catppuccin.json"),
	)
	if err != nil {
		return styles.ErrorStyle.Render("Renderer error: " + err.Error())
	}

	rendered, err := renderer.Render(string(data))
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
