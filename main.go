package main

import (
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func bumpVersion() {
	file, err := os.Create("version.txt")

	if err != nil {
		return
	}
	n, _ := strconv.Atoi(version)
	file.WriteString(fmt.Sprintf("%d", n+1))

	file.Close()
}

//go:embed version.txt
var version string

func main() {
	defer func() {
		bumpVersion()
	}()
	m := NewModel()

	tea.LogToFile(fmt.Sprintf("default_%s.log", version), "")

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// Model: app state
type Model struct {
}

// NewModel: initial model
func NewModel() Model {

	return Model{}
}

// Init: kick off the event loop
func (m Model) Init() tea.Cmd {
	slog.Info("Model Init Called")
	return nil
}

// Update: handle Msgs
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		return m, cmd
	}
	return m, nil
}

// View: return a string based on the state of our model
func (m Model) View() string {
	s := "Hello World"

	return s
}
