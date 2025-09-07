package app

import (
	"github.com/antoniosarro/reltrace/internal/config"
	"github.com/antoniosarro/reltrace/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the main application
type App struct {
	config config.AppConfig
	ui     *ui.Model
}

// New creates a new application instance
func New() *App {
	appConfig := config.DefaultConfig()

	return &App{
		config: appConfig,
		ui:     ui.New(appConfig),
	}
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	return a.ui.Init()
}

// Update handles application updates
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.ui, cmd = a.ui.Update(msg)
	return a, cmd
}

// View renders the application
func (a *App) View() string {
	return a.ui.View()
}
