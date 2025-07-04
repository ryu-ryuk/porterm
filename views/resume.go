package views

import (
	"strings"

	"porterm/internal/data"
	"porterm/styles"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var rawResumeBodyContent string
var resumeContactInfo string

func init() {
	lines := strings.Split(data.ResumeData, "\n")
	if len(lines) < 2 {
		rawResumeBodyContent = styles.ErrorStyle.Render("Resume content too short or malformed.")
		resumeContactInfo = ""
		return
	}

	resumeContactInfo = lines[1]
	rawResumeBodyContent = strings.Join(lines[2:], "\n")
}

func GetResumeContactInfo() string {
	return resumeContactInfo
}

// About glamour-catppuccin.json
// Getting used by the Glamour library for rendering Markdown documents.
// This file defines a color scheme based on the Catppuccin theme, specifically the "Mocha" variant.
// The colors are chosen to match the Catppuccin Mocha palette, providing a consistent and visually appealing look for Markdown content.
// The color values are in hexadecimal format, and the properties define how different elements of the Markdown document should be styled.
// The properties include colors for the document background, headings, paragraphs, blockquotes, lists, code blocks, emphasized text, strong text, and links.
// This JSON structure can be used in applications that support custom Markdown rendering styles, allowing users to enjoy a cohesive and aesthetically pleasing reading experience.
// Catppuccin Mocha Theme for Glamour Markdown Renderer

// prepares the styled resume content for the viewport using glamour.
func ResumeContentForViewport(width int) string {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(data.GlamourStyle),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return styles.ErrorStyle.Render("Renderer error: " + err.Error())
	}

	rendered, err := renderer.Render(rawResumeBodyContent)
	if err != nil {
		return styles.ErrorStyle.Render("Markdown error: " + err.Error())
	}

	return strings.TrimSpace(rendered)
}

func Resume(width int) string {
	if rawResumeBodyContent == "" {
		return lipgloss.JoinVertical(lipgloss.Left,
			styles.Heading.Render("Resume"),
			styles.ErrorStyle.Render("Resume content not loaded."),
			"",
			styles.Help.Render("Press esc to return to menu."),
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.Heading.Render("Resume"),
		styles.ContactInfo.Render(resumeContactInfo),
		"",
		ResumeContentForViewport(width),
		"",
		styles.Help.Render("Press esc to return to menu."),
	)
}
