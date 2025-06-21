package main

import (
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"terminal-portfolio/model"
)

func main() {
	p := tea.NewProgram(model.New())
	if _, err := p.Run(); err != nil {
		println("Error running program:", err)
		os.Exit(1)
	}
}

