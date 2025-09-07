package ui

import (
	"github.com/antoniosarro/reltrace/internal/ui/components/database"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the UI model
func (m *Model) Init() tea.Cmd {
	return m.dbSelector.Init()
}

// Update handles UI updates
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Handle window size changes
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case database.DatabaseSelectedMsg:
		// Update the config form with selected database type
		m.configForm.SetDatabaseType(msg.DatabaseType)
		m.state = ConfigurationView
		return m, m.configForm.Focus()
	}

	// Update current view
	switch m.state {
	case DatabaseSelectionView:
		return m.updateDatabaseSelector(msg)
	case ConfigurationView:
		return m.updateConfigForm(msg)
	}

	return m, nil
}

// View renders the UI
func (m *Model) View() string {
	switch m.state {
	case DatabaseSelectionView:
		return m.dbSelector.View()
	case ConfigurationView:
		return m.configForm.View()
	}
	return ""
}
