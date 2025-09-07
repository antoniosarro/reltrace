package main

import (
	"log"
	"os"

	"github.com/antoniosarro/reltrace/internal/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := app.New()

	program := tea.NewProgram(
		app,
		tea.WithAltScreen(),
	)

	if _, err := program.Run(); err != nil {
		log.Printf("error running application: %v", err)
		os.Exit(1)
	}
}
