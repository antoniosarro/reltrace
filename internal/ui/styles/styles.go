package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles contains all the styling definitions
type Styles struct {
	// Base styles
	Focused lipgloss.Style
	Blurred lipgloss.Style
	Cursor  lipgloss.Style
	Help    lipgloss.Style

	// UI elements
	Title    lipgloss.Style
	Status   lipgloss.Style
	Progress lipgloss.Style
	Button   lipgloss.Style

	// Database type styles
	MySQL      lipgloss.Style
	PostgreSQL lipgloss.Style
	SQLite     lipgloss.Style

	// State styles
	Success lipgloss.Style
	Error   lipgloss.Style
	Warning lipgloss.Style
	Info    lipgloss.Style
}

// New creates a new styles instance
func New() *Styles {
	return &Styles{
		// Base styles
		Focused: lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
		Blurred: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		Cursor:  lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
		Help:    lipgloss.NewStyle().Foreground(lipgloss.Color("244")),

		// UI elements
		Title: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1),

		Status: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1),

		Progress: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Padding(1, 2),

		Button: lipgloss.NewStyle().
			Background(lipgloss.Color("#6366F1")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			MarginRight(1),

		// Database type styles
		MySQL: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E97627")).
			Bold(true),

		PostgreSQL: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#336791")).
			Bold(true),

		SQLite: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#003B57")).
			Bold(true),

		// State styles
		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00C851")).
			Bold(true),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4444")).
			Bold(true),

		Warning: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFBB33")).
			Bold(true),

		Info: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#33B5E5")),
	}
}

// GetDatabaseStyle returns the appropriate style for a database type
func (s *Styles) GetDatabaseStyle(dbType string) lipgloss.Style {
	switch dbType {
	case "mysql":
		return s.MySQL
	case "postgresql":
		return s.PostgreSQL
	case "sqlite3":
		return s.SQLite
	default:
		return s.Info
	}
}
