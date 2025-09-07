package configs

import (
	"github.com/antoniosarro/reltrace/internal/database/models"
	tea "github.com/charmbracelet/bubbletea"
)

// View renders the configuration form
func (c *ConfigForm) View() string {

	switch c.step {
	case 0:
		return c.viewDatabaseConfig()
	case 1:
		return c.viewModeSelection()
	case 2:
		return c.viewTargetSelection()
	default:
		return c.viewDatabaseConfig()
	}
}

// Focus focuses the form
func (c *ConfigForm) Focus() tea.Cmd {
	c.focusIndex = 0
	return c.inputs[0].Focus()
}

// Update handles configuration form updates
func (c *ConfigForm) Update(msg tea.Msg) (*ConfigForm, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			return c.handleNavigation(msg.String())
		case "1", "2", "3", "4":
			// Mode selection
			switch c.step {
			case 1:
				mode := models.DumpMode(int(msg.String()[0] - '1'))
				if mode >= models.StructureOnly && mode <= models.StructureAndDataIncludingOnly {
					// Move to target selection
					c.mode = mode
					c.step = 2
					return c, nil
				}
			case 2: // Target selection
				if msg.String() == "1" {
					c.target = models.ToFile
				} else if msg.String() == "2" {
					c.target = models.ToDatabase
				}
				if c.target == models.ToFile || c.validateTargetDbConfig() {
					return c, c.submitConfig()
				}
			}
		case "backspace":
			if c.step > 0 {
				c.step--
				return c, nil
			}
		}
	}

	// Update inputs only in step 0
	if c.step == 0 {
		cmds := make([]tea.Cmd, len(c.inputs))
		for i := range c.inputs {
			c.inputs[i], cmds[i] = c.inputs[i].Update(msg)
		}
		return c, tea.Batch(cmds...)
	}

	return c, nil
}

// SetDatabaseType sets the database type and updates placeholders
func (c *ConfigForm) SetDatabaseType(dbType models.DatabaseType) {
	c.dbType = dbType
	c.updatePlaceholders()
}

// updatePlaceholders updates input placeholders based on database type
func (c *ConfigForm) updatePlaceholders() {
	switch c.dbType {
	case models.MySQL:
		c.inputs[0].Placeholder = "Host (e.g., localhost)"
		c.inputs[1].Placeholder = "Port (default: 3306)"
		c.inputs[4].Placeholder = "Database Name"
	case models.PostgreSQL:
		c.inputs[0].Placeholder = "Host (e.g., localhost)"
		c.inputs[1].Placeholder = "Port (default: 5432)"
		c.inputs[4].Placeholder = "Database Name"
	case models.SQLite3:
		c.inputs[0].Placeholder = "Not needed for SQLite"
		c.inputs[1].Placeholder = "Not needed for SQLite"
		c.inputs[2].Placeholder = "Not needed for SQLite"
		c.inputs[3].Placeholder = "Not needed for SQLite"
		c.inputs[4].Placeholder = "SQLite File Path (e.g., ./database.db)"
	}
}
