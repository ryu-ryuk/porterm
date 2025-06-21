package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"porterm/model"
	"os"
)

func main() {
	p := tea.NewProgram(model.New())
	if _, err := p.Run(); err != nil {
		println("Error running program:", err)
		os.Exit(1)
	}
}
