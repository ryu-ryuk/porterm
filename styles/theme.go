package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// catp colors
const (
	Rosewater = lipgloss.Color("#f5e0dc")
	Flamingo  = lipgloss.Color("#f2cdcd")
	Pink      = lipgloss.Color("#f5c2e7")
	Mauve     = lipgloss.Color("#cba6f7")
	Red       = lipgloss.Color("#f38ba8")
	Maroon    = lipgloss.Color("#eba0ac")
	Peach     = lipgloss.Color("#fab387")
	Yellow    = lipgloss.Color("#f9e2af")
	Green     = lipgloss.Color("#a6e3a1")
	Teal      = lipgloss.Color("#94e2d5")
	Sky       = lipgloss.Color("#89dceb")
	Sapphire  = lipgloss.Color("#74c7ec")
	Blue      = lipgloss.Color("#89b4fa")
	Lavender  = lipgloss.Color("#b4befe")
	Text      = lipgloss.Color("#cdd6f4")
	Subtext1  = lipgloss.Color("#bac2de")
	Subtext0  = lipgloss.Color("#a6adc8")
	Overlay2  = lipgloss.Color("#9399b2")
	Overlay1  = lipgloss.Color("#7f849c")
	Overlay0  = lipgloss.Color("#6c7086")
	Surface2  = lipgloss.Color("#585b70")
	Surface1  = lipgloss.Color("#45475a")
	Surface0  = lipgloss.Color("#313244")
	Base      = lipgloss.Color("#1e1e2e")
	Mantle    = lipgloss.Color("#181825")
	Crust     = lipgloss.Color("#11111b")
)

var (
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Background(Base).
			Foreground(Text)

	MenuItem = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(Text)

	SelectedMenuItem = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(Green).
				Background(Surface0).
				Bold(true)

	Heading = lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder(), false, false, true, false).
		BorderForeground(Mauve).
		Foreground(Mauve).
		Padding(0, 1).
		MarginBottom(1)

	SubHeading = lipgloss.NewStyle().
			Foreground(Green).
			Bold(true).
			MarginTop(1).
			PaddingLeft(1)

	Content = lipgloss.NewStyle().
		Foreground(Subtext1).
		PaddingLeft(2)

	Emphasis = lipgloss.NewStyle().
			Foreground(Blue).
			Bold(true)

	Help = lipgloss.NewStyle().
		Foreground(Overlay0).
		PaddingTop(1)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Red).
			Bold(true)

	ResumeContainer = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Sapphire).
			Padding(1, 3).
			Align(lipgloss.Left)

	ContactInfo = lipgloss.NewStyle().
			Foreground(Subtext0).
			PaddingBottom(1)

	ViewportStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(Surface2)

	//  using JSON file instead of the glamour styling functions
)

// render bordered block
func BorderedBlock(content string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Green).
		Padding(1, 2).
		Width(lipgloss.Width(content) + 4).
		Render(content)
}

// render menu item
func RenderMenuItem(number string, text string) string {
	return MenuItem.Render(Emphasis.Render(number+".") + " " + text)
}
