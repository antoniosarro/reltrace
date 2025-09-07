package configs

import (
	"github.com/antoniosarro/reltrace/internal/database/models"
	tea "github.com/charmbracelet/bubbletea"
)

// handleNavigation handles navigation between form fields
func (c *ConfigForm) handleNavigation(key string) (*ConfigForm, tea.Cmd) {
	if c.step != 0 {
		return c, nil
	}

	switch key {
	case "enter":
		if c.focusIndex == len(c.inputs) {
			if c.validateInputs() {
				c.step = 1 // Move to mode selection
				return c, nil
			}
		}
	case "up", "shift+tab":
		c.focusIndex--
	case "down", "tab":
		c.focusIndex++
	}

	maxIndex := len(c.inputs)
	if c.focusIndex > maxIndex {
		c.focusIndex = 0
	} else if c.focusIndex < 0 {
		c.focusIndex = maxIndex
	}

	cmds := make([]tea.Cmd, len(c.inputs))
	for i := 0; i < len(c.inputs); i++ {
		if i == c.focusIndex {
			cmds[i] = c.inputs[i].Focus()
			c.inputs[i].PromptStyle = c.styles.Focused
			c.inputs[i].TextStyle = c.styles.Focused
		} else {
			c.inputs[i].Blur()
			c.inputs[i].PromptStyle = c.styles.Blurred
			c.inputs[i].TextStyle = c.styles.Blurred
		}
	}
	return c, tea.Batch(cmds...)
}

// validateInputs validates the form inputs
func (c *ConfigForm) validateInputs() bool {
	switch c.dbType {
	case models.SQLite3:
		return c.inputs[4].Value() != "" // File path required
	default:
		required := []int{0, 2, 4} // Host, username, database
		for _, i := range required {
			if c.inputs[i].Value() == "" {
				return false
			}
		}

		// Root table/key required for specific modes
		if c.mode == models.StructureAndDataExcluding || c.mode == models.StructureAndDataIncludingOnly {
			if c.inputs[5].Value() == "" || c.inputs[6].Value() == "" {
				return false
			}
		}

		return true
	}
}

// validateTargetDbConfig validates target database configuration
func (c *ConfigForm) validateTargetDbConfig() bool {
	// For now, assume target DB config is valid
	// In a full implementation, collect target DB details
	return true
}

// submitConfig creates and submits the final configuration
func (c *ConfigForm) submitConfig() tea.Cmd {
	config := c.buildConfig()
	return func() tea.Msg {
		return ConfigCompletedMsg{Config: config}
	}
}

// Rest of the methods remain similar but updated for new config structure...
func (c *ConfigForm) buildConfig() models.DumpConfig {
	sourceConfig := models.DatabaseConfig{
		Type:     c.dbType,
		Host:     c.inputs[0].Value(),
		Port:     c.inputs[1].Value(),
		User:     c.inputs[2].Value(),
		Password: c.inputs[3].Value(),
	}

	if c.dbType == models.SQLite3 {
		sourceConfig.FilePath = c.inputs[4].Value()
	} else {
		sourceConfig.Database = c.inputs[4].Value()
		// Set default ports
		if sourceConfig.Port == "" {
			switch c.dbType {
			case models.MySQL:
				sourceConfig.Port = "3306"
			case models.PostgreSQL:
				sourceConfig.Port = "5432"
			}
		}
	}

	config := models.DumpConfig{
		SourceConfig: sourceConfig,
		Mode:         c.mode,
		Target:       c.target,
		OutputPath:   c.inputs[7].Value(),
	}

	// Add root table/key for modes that need them
	if c.mode == models.StructureAndDataExcluding || c.mode == models.StructureAndDataIncludingOnly {
		config.RootTable = c.inputs[5].Value()
		config.RootPrimaryKey = c.inputs[6].Value()
	}

	return config
}
