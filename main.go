package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Info("START")

	m := model{}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Error("Error while starting program", err)
		os.Exit(1)
	}
}

type model struct {
	viewportWidth  int
	viewportHeight int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyRunes:
			if msg.String() == "q" {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		log.Infof("Height: %v Width: %v", msg.Height, msg.Width)
		m.viewportHeight = msg.Height
		m.viewportWidth = msg.Width
	}

	return m, nil
}

func (m model) View() string {
	sqr := lipgloss.NewStyle().Width(40).Height(16).Border(lipgloss.RoundedBorder()).Render("snake-shell")

	return lipgloss.Place(m.viewportWidth, m.viewportHeight, lipgloss.Center, lipgloss.Center, sqr)
}
