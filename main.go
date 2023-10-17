package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/chancesm/tui-exploration/appmodel"
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

	tea.LogToFile(fmt.Sprintf("default_%s.log", version), "")

	m := appmodel.NewModel()

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
	}
}
