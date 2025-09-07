package database

import (
	"fmt"
	"strings"

	"github.com/antoniosarro/reltrace/internal/database/models"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the database selector
func (d *DatabaseSelector) Init() tea.Cmd {
	return nil
}

// Update handles database selector updates
func (d *DatabaseSelector) Update(msg tea.Msg) (*DatabaseSelector, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			d.selected--
			if d.selected < 0 {
				d.selected = len(d.choices) - 1
			}
		case "down", "j":
			d.selected++
			if d.selected >= len(d.choices) {
				d.selected = 0
			}
		case "enter":
			return d, func() tea.Msg {
				return DatabaseSelectedMsg{DatabaseType: d.choices[d.selected]}
			}
		}
	}

	return d, nil
}

// View renders the database selector
func (d *DatabaseSelector) View() string {
	var b strings.Builder

	b.WriteString(d.styles.Title.Render("🗃️  Reltrace - Multi-Database Recursive Backup"))
	b.WriteString("\n\n")
	b.WriteString("Select your database type:\n\n")

	// Find the maximum length of dbType strings for alignment
	maxDBTypeLength := 0
	for _, dbType := range d.choices {
		if len(string(dbType)) > maxDBTypeLength {
			maxDBTypeLength = len(string(dbType))
		}
	}

	for i, dbType := range d.choices {
		cursor := "  "
		if i == d.selected {
			cursor = "> "
		}
		var icon, desc string
		style := d.styles.Blurred
		if i == d.selected {
			style = d.styles.Focused
		}

		switch dbType {
		case models.MySQL:
			icon = "🐬"
			desc = "MySQL Database"
		case models.PostgreSQL:
			icon = "🐘"
			desc = "PostgreSQL Database"
		case models.SQLite3:
			icon = "📁"
			desc = "SQLite3 Database"
		}

		// Calculate padding needed to align descriptions
		dbTypeStr := style.Render(string(dbType))
		padding := strings.Repeat(" ", maxDBTypeLength-len(string(dbType))+2)

		b.WriteString(fmt.Sprintf("%s%s %s%s %s\n",
			cursor, icon, dbTypeStr, padding, desc))
	}

	b.WriteString("\n" + d.styles.Help.Render("• Use ↑/↓ or j/k to navigate • Enter to select • Ctrl+C to quit"))
	return b.String()
}
