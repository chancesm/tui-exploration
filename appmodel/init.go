package appmodel

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	slog.Info("Model Init Called")
	return nil
}
