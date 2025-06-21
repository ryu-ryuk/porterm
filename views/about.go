package views

import (
	"github.com/charmbracelet/lipgloss"
	"porterm/styles"
)

var funFacts = []string{
    `"waking up everyday to wanting to go back to sleep"`,
    `"part-time cat, full-time snack connoisseur."`,
	`"can never debug my life, maybe i need a debugger for that ;) "`,
	`"i'm not lazy!!!"`,
	`"i put the 'pro' in procrastination"`,
	`"if at first you don't succeed, redefine success."`,
	`"some parts are missing..."`,
	`"i used to be indecisive, but now i'm not so sure."`,
	`"my life feels like a test i didn't study for."`,
	`"if you think nobody cares if you're alive, try missing a couple of payments."`,
	`"i'm 6.9ft tall!"`,
}

func About(funFactIndex int) string {
	// styles for this view
	highlight := lipgloss.NewStyle().
		Foreground(styles.Peach).
		Bold(true)
	index := funFactIndex % len(funFacts)
	funFactText := funFacts[index]

	techStack := lipgloss.NewStyle().
		Foreground(styles.Lavender).
		MarginLeft(2)

	funFact := lipgloss.NewStyle().
		Foreground(styles.Yellow).
		Italic(true)

	quote := lipgloss.NewStyle().
		Foreground(styles.Mauve).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Surface1).
		Padding(1, 2).
		MarginTop(1)

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.Heading.Render("About Me"),
		
		lipgloss.JoinHorizontal(lipgloss.Top,
			styles.Content.Render("hola! I'm"),
			styles.Content.Render("Alok Ranjan - "),
			highlight.Render("Backend Developer @ IITM Paradox"),
		),
		
		"\n",
		styles.SubHeading.Render("Tech Stack:"),
		techStack.Render("• Rust: Systems programming, security tools, TUIs"),
		techStack.Render("• Go: Backend services, CLIs"),
		techStack.Render("• Python: Automation, scripting, AI/ML"),
		techStack.Render("• Linux: Advanced system customization"),
		// funFact.Render(`"turning 'I followed the recipe exactly' into a culinary disaster"`),

		"\n",
		styles.SubHeading.Render("Passions:"),
		styles.Content.Render("(´∀`) Learning "), highlight.Render("	Cybersecurity"),
		styles.Content.Render("⊂(・ヮ・⊂) Creating "), highlight.Render("	terminal-based xperiences"),
		styles.Content.Render("(｀∇´ゞ Designing "), highlight.Render("	visually rich TUIs | sneaking catppuccin into all of it ;) "),
		styles.Content.Render("(ノAヽ) Searching for "), highlight.Render("	the will to live!"),

		funFact.Render(),
		
		"\n",
		quote.Render(funFactText),
		
		"\n\n",
		styles.Help.Render("Press esc to return to menu"),
	)
}
