package views

import (
	"github.com/charmbracelet/lipgloss"
	"porterm/styles"
)

func Badges() string {
	// badge styling
	badgeStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		Margin(0, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Surface1)

	webringStyle := lipgloss.NewStyle().
		Foreground(styles.Green).
		Padding(0, 1)

	arrowStyle := lipgloss.NewStyle().
		Foreground(styles.Peach)

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.Heading.Render("Achievements & Webrings"),
		
		// webring badges
		styles.SubHeading.Render("Webrings"),
		styles.SubHeading.Render("--------"),
		lipgloss.JoinHorizontal(lipgloss.Top,
			arrowStyle.Render("←"),
			badgeStyle.Background(styles.Surface0).Render("WebOps Webring"),
			arrowStyle.Render("→"),
		),
		webringStyle.Render("https://iitm-paradox.github.io/webring/"),
		"",
		
		lipgloss.JoinHorizontal(lipgloss.Top,
			arrowStyle.Render("←"),
			badgeStyle.Background(styles.Mauve).Render("Catppuccin Webring"),
			arrowStyle.Render("→"),
		),
		webringStyle.Render("https://ctp-webr.ing/"),
		"",
		
		// badges section
		styles.SubHeading.Render("Certifications"),
		styles.SubHeading.Render("--------"),

		badgeStyle.Background(styles.Blue).Render("Google IT Support Professional"),
		badgeStyle.Background(styles.Green).Render("Generative AI Course"),
		badgeStyle.Background(styles.Yellow).Render("SAWIT.AI Learnathon"),
		badgeStyle.Background(styles.Teal).Render("30+ Google Cloud Skill Badges"),
		"",
		
		styles.Help.Render("Press esc to return to menu"),
	)
}
