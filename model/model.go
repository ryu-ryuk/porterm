package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"time"
	"strings"
	"math"
	"terminal-portfolio/styles"
	"terminal-portfolio/views"
)

type Model struct {
	view          string
	viewport      viewport.Model
	ready         bool
	zoomLevel     int
	minZoomWidth  int
	maxZoomWidth  int
	width         int
	height        int
	animationFrame int
	bannerColors  []lipgloss.Color
}

const (
	defaultResumeWidth = 80
	zoomStep           = 10
)

func New() Model {
	return Model{
		view:         "menu",
		zoomLevel:    0,
		minZoomWidth: 40,
		maxZoomWidth: 160,
		bannerColors: []lipgloss.Color{
			styles.Mauve,
			styles.Peach,
			styles.Green,
			styles.Sky,
			styles.Lavender,
			styles.Yellow,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
			return tickMsg{}
		}),
	)
}

type tickMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tickMsg:
		m.animationFrame++
		return m, tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
			return tickMsg{}
		})

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		hPadding := styles.AppStyle.GetHorizontalPadding() +
			styles.ResumeContainer.GetHorizontalPadding() +
			styles.ResumeContainer.GetHorizontalBorderSize()

		vPadding := styles.AppStyle.GetVerticalPadding() +
			styles.ResumeContainer.GetVerticalPadding() +
			styles.ResumeContainer.GetVerticalBorderSize() +
			lipgloss.Height(styles.Heading.Render("")) +
			lipgloss.Height(styles.ContactInfo.Render("")) +
			lipgloss.Height(styles.Help.Render("")) + 2

		if !m.ready {
			m.viewport = viewport.New(msg.Width-hPadding, msg.Height-vPadding)
			m.viewport.Style = styles.ViewportStyle
			m.ready = true
		} else {
			m.viewport.Width = msg.Width - hPadding
			m.viewport.Height = msg.Height - vPadding
		}

		m.maxZoomWidth = msg.Width - hPadding

		if m.view == "resume" || !m.ready {
			m.updateResumeContent()
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.view = "about"
		case "2":
			m.view = "projects"
		case "3":
			m.view = "resume"
			if m.ready {
				m.updateResumeContent()
			}
		case "4":
			m.view = "badges"

		case "+", "=":
			if m.view == "resume" {
				m.zoomLevel++
				m.updateResumeContent()
				return m, nil
			}
		case "-":
			if m.view == "resume" {
				m.zoomLevel--
				m.updateResumeContent()
				return m, nil
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.view = "menu"
			if m.ready {
				m.viewport.SetContent("")
			}
		}
	}

	if m.view == "resume" {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) updateResumeContent() {
	baseContentWidth := defaultResumeWidth
	currentContentWidth := baseContentWidth + (m.zoomLevel * zoomStep)

	if currentContentWidth > m.viewport.Width {
		currentContentWidth = m.viewport.Width
	}
	if currentContentWidth < m.minZoomWidth {
		currentContentWidth = m.minZoomWidth
	}
	if currentContentWidth > m.maxZoomWidth {
		currentContentWidth = m.maxZoomWidth
	}

	styles.ResumeContainer = styles.ResumeContainer.Copy().
		Width(currentContentWidth +
			styles.ResumeContainer.GetHorizontalPadding() +
			styles.ResumeContainer.GetHorizontalBorderSize())

	m.viewport.SetContent(views.ResumeContentForViewport(currentContentWidth))
	m.viewport.GotoTop()
}


func glitchEffect(frame int, text string) string {
	lines := strings.Split(text, "\n")
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	result := make([]string, len(lines))
	for i, line := range lines {
		newLine := make([]rune, maxWidth)
		// init with spaces
		for j := range newLine {
			newLine[j] = ' '
		}

		// wave effect to each character
		for pos, char := range line {
			waveOffset := int(1.5 * math.Sin(float64(pos+frame)/3.0))
			newPos := pos + waveOffset
			if newPos >= 0 && newPos < maxWidth {
				newLine[newPos] = char
			}
		}
		result[i] = string(newLine)
	}
	return strings.Join(result, "\n")
}
func (m Model) colorizeVectorArt() string {

	mochaColors := []lipgloss.Color{
		styles.Rosewater, // #f5e0dc
		styles.Flamingo,  // #f2cdcd
		styles.Pink,      // #f5c2e7
		styles.Mauve,     // #cba6f7
		styles.Red,       // #f38ba8
		styles.Maroon,    // #eba0ac
		styles.Peach,     // #fab387
		styles.Yellow,    // #f9e2af
		styles.Green,     // #a6e3a1
		styles.Teal,      // #94e2d5
		styles.Sky,       // #89dceb
		styles.Sapphire,  // #74c7ec
		styles.Blue,      // #89b4fa
		styles.Lavender,  // #b4befe
		styles.Text,      // #cdd6f4
		styles.Subtext1,  // #bac2de
		styles.Subtext0,  // #a6adc8
		styles.Overlay2,  // #9399b2
		styles.Overlay1,  // #7f849c
		styles.Overlay0,  // #6c7086
		styles.Surface2,  // #585b70
		styles.Surface1,  // #45475a
		styles.Surface0,  // #313244
		styles.Base,      // #1e1e2e
		styles.Mantle,    // #181825
		styles.Crust,     // #11111b
	}

	vectorArt := []string{
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀",
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⡋⠁⠀⠀⠀⠀⢀⣀⣀⡀",
		"⠀⠀⠀⠀⠀⠠⠒⣶⣶⣿⣿⣷⣾⣿⣿⣿⣿⣛⣋⣉⠀⠀",
		"⠀⠀⠀⠀⢀⣤⣞⣫⣿⣿⣿⡻⢿⣿⣿⣿⣿⣿⣦⡀⠀⠀",
		"⠀⠀⣶⣾⡿⠿⠿⠿⠿⠋⠈⠀⣸⣿⣿⣿⣿⣷⡈⠙⢆⠀",
		"⠀⠀⠉⠁⠀⠤⣤⣤⣤⣤⣶⣾⣿⣿⣿⣿⠿⣿⣷⠀⠀⠀",
		"⠀⠀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⡿⠟⠁⠀⢹⣿⠀⠀⠀",
		"⢠⣾⣿⣿⣿⣿⠟⠋⠉⠛⠋⠉⠁⣀⠀⠀⠀⠸⠃⠀⠀⠀",
		"⣿⣿⣿⣿⠹⣇⠀⠀⠀⠀⢀⡀⠀⢀⡙⢷⣦⣄⡀⠀⠀⠀",
		"⣿⢿⣿⣿⣷⣦⠤⠤⠀⠀⣠⣿⣶⣶⣿⣿⣿⣿⣿⣷⣄⠀",
		"⠈⠈⣿⡿⢿⣿⣿⣷⣿⣿⡿⢿⣿⣿⣁⡀⠀⠀⠉⢻⣿⣧",
		"⠀⢀⡟⠀⠀⠉⠛⠙⠻⢿⣦⡀⠙⠛⠯⠤⠄⠀⠀⠈⠈⣿",
		"⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⡆⠀⠀⠀⠀⠀⠀⠀⢀⠟",
		"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	}

	// coloring each line and cycling it through the mocha colors
	coloredLines := make([]string, len(vectorArt))
	for i, line := range vectorArt {
		colorIndex := (i + m.animationFrame) % len(mochaColors)
		style := lipgloss.NewStyle().Foreground(mochaColors[colorIndex])
		coloredLines[i] = style.Render(line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, coloredLines...)
}

func (m Model) View() string {

	// to center all content vertically and horizontally
	centerStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center)

	switch m.view {
	case "about":
		content := views.About()
		return centerStyle.Render(styles.AppStyle.Render(content))

	case "projects":
		content := views.Projects(m.width)
		return centerStyle.Render(styles.AppStyle.Render(content))
		
	case "badges":
		content := views.Badges()
		return centerStyle.Render(styles.AppStyle.Render(content))
		
	case "resume":
		if !m.ready {
			return centerStyle.Render("im doing my best...")
		}

		renderedResume := lipgloss.JoinVertical(lipgloss.Left,
			styles.Heading.Render("Resume"),
			styles.ContactInfo.Render(views.GetResumeContactInfo()),
			m.viewport.View(),
			styles.Help.Render("Scroll: ↑↓ / PgUp/PgDn | Zoom: + / - | Back: Esc | Quit: q"),
		)

		return centerStyle.Render(styles.AppStyle.Render(
			styles.ResumeContainer.Render(renderedResume),
		))
	default:
		vectorBlock := m.colorizeVectorArt()
		vectorStyle := lipgloss.NewStyle().Align(lipgloss.Center)
		vectorArtBlock := vectorStyle.Render(vectorBlock)


		bannerTop := 
		`
 ___ _/ /__  / /__  _______ ____    (_)__ ____ 
/ _  / / _ \/   _/ / __/ _  / _ \  / / _  / _ \
\_,_/_/\___/_/\_\ /_/  \_,_/_//_/_/ /\_,_/_//_/
                               |___/           

		`
		bannerTopStyle := lipgloss.NewStyle().
			Foreground(styles.Lavender).
			Bold(true).
			Align(lipgloss.Center)

		baseKanji := 
`
░▀▀░▀▀░▀▀░▀▀
`
		
		kanjiPart := glitchEffect(m.animationFrame, baseKanji)
		kanjiStyle := lipgloss.NewStyle().
			Foreground(styles.Mauve).
			Bold(true).
			Align(lipgloss.Center)

		// join banner parts
		fullBanner := vectorArtBlock + "\n" +
			bannerTopStyle.Render(bannerTop) + "\n" +
			kanjiStyle.Render(kanjiPart)

		// animated banner color
		// colorIndex := m.animationFrame % len(m.bannerColors)
		// currentColor := m.bannerColors[colorIndex]
		// bannerStyle := lipgloss.NewStyle().
		// 	Foreground(currentColor).
		// 	Bold(true).
		// 	MarginBottom(1).
		// 	Align(lipgloss.Center)

		// animated border for menu
		borderColors := []lipgloss.Color{
			styles.Sapphire,
			styles.Blue,
			styles.Lavender,
			styles.Mauve,
		}
		borderColor := borderColors[m.animationFrame%len(borderColors)]

		menuBorder := lipgloss.Border{
			Top:         "─",
			Bottom:      "─",
			Left:        "│",
			Right:       "│",
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "╰",
			BottomRight: "╯",
		}
		menuStyle := lipgloss.NewStyle().
			Border(menuBorder).
			BorderForeground(borderColor).
			Padding(1, 3).
			MarginBottom(1).
			Align(lipgloss.Center)

		// catp colors
		menuColors := []lipgloss.Color{
			styles.Peach,
			styles.Sky,
			styles.Lavender,
			styles.Green,
		}
		menuItems := lipgloss.JoinVertical(lipgloss.Left,
			lipgloss.NewStyle().Foreground(menuColors[0]).Render(styles.RenderMenuItem("1", "Bout Me")),
			lipgloss.NewStyle().Foreground(menuColors[1]).Render(styles.RenderMenuItem("2", "My Works")),
			lipgloss.NewStyle().Foreground(menuColors[2]).Render(styles.RenderMenuItem("3", "Resume")),
			lipgloss.NewStyle().Foreground(menuColors[3]).Render(styles.RenderMenuItem("4", "Webrings & Badges")),
		)

		// hehe
		footerQuotes := []string{
			"sudo make life awesome",
			"git commit -m 'another day, another struggle'",
			"while(!succeed) { tryAgain(); }",
			"echo 'Hello World!' > /dev/life",
			"catppuccin > rose-pine > nord (fight me)",
			"diagramming in Mermaid, themed in Mocha",
			"terminal UIs are peak dev aesthetics",
		}
		footerIndex := int(time.Now().Unix()) % len(footerQuotes)
		footer := lipgloss.NewStyle().
			Foreground(styles.Overlay1).
			Italic(true).
			MarginTop(1).
			Render(footerQuotes[footerIndex])

		// Avenger$ assemble !! 
		content := lipgloss.JoinVertical(lipgloss.Center,
			fullBanner,
			menuStyle.Render(menuItems),
			footer,
			"",
			styles.Help.Render("Press 1/2/3/4 to navigate | q to quit"),
		)

		return centerStyle.Render(styles.AppStyle.Render(content))
	}
}