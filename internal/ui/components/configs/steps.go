package configs

import (
	"fmt"
	"strings"

	"github.com/antoniosarro/reltrace/internal/database/models"
)

// viewDatabaseConfig renders the database configuration step
func (c *ConfigForm) viewDatabaseConfig() string {
	var b strings.Builder

	dbStyle := c.styles.GetDatabaseStyle(string(c.dbType))
	title := fmt.Sprintf("Configure %s Connection", dbStyle.Render(strings.ToUpper(string(c.dbType))))
	b.WriteString(c.styles.Title.Render(title))
	b.WriteString("\n\n")

	labels := []string{
		"Host:", "Port:", "Username:", "Password:",
		"Database/File:", "Root Table:", "Primary Key:", "Output Path:",
	}

	for i := range c.inputs {
		// Skip irrelevant fields for SQLite
		if c.dbType == models.SQLite3 && i < 4 {
			continue
		}

		// Skip root table/key for structure-only dumps
		if (i == 5 || i == 6) && c.mode == models.StructureOnly {
			continue
		}

		b.WriteString(labels[i] + "\n")
		b.WriteString(c.inputs[i].View())
		b.WriteString("\n\n")
	}

	button := "[ Next ]"
	if c.focusIndex == len(c.inputs) {
		button = c.styles.Button.Render("[ Next ]")
	} else {
		button = c.styles.Blurred.Render("[ Next ]")
	}
	b.WriteString(button + "\n\n")

	b.WriteString(c.styles.Help.Render("• Tab to navigate • Enter to continue • Ctrl+C to quit"))
	return b.String()
}

// viewModeSelection renders the mode selection step
func (c *ConfigForm) viewModeSelection() string {
	var b strings.Builder

	b.WriteString(c.styles.Title.Render("Select Export Mode"))
	b.WriteString("\n\n")

	modes := []struct {
		mode models.DumpMode
		name string
		desc string
	}{
		{models.StructureOnly, "Structure Only", "Export database structure without data"},
		{models.StructureAndData, "Structure + All Data", "Export complete database structure and all data"},
		{models.StructureAndDataExcluding, "Structure + Data (Excluding)", "Export all data except records related to specified root"},
		{models.StructureAndDataIncludingOnly, "Structure + Data (Including Only)", "Export only records related to specified root"},
	}

	for i, m := range modes {
		prefix := fmt.Sprintf("%d. ", i+1)
		if c.mode == m.mode {
			b.WriteString(c.styles.Focused.Render(prefix + m.name))
		} else {
			b.WriteString(c.styles.Blurred.Render(prefix + m.name))
		}
		b.WriteString(" - " + m.desc + "\n")
	}

	b.WriteString("\n")
	b.WriteString(c.styles.Help.Render("• Press 1-4 to select mode • Backspace to go back"))
	return b.String()
}

// viewTargetSelection renders the target selection step
func (c *ConfigForm) viewTargetSelection() string {
	var b strings.Builder

	b.WriteString(c.styles.Title.Render("Select Export Target"))
	b.WriteString("\n\n")

	b.WriteString("Selected Mode: ")
	b.WriteString(c.styles.Info.Render(c.mode.String()))
	b.WriteString("\n\n")

	targets := []struct {
		target models.DumpTarget
		name   string
		desc   string
	}{
		{models.ToFile, "Export to File", "Save dump to SQL file"},
		{models.ToDatabase, "Export to Database", "Import directly to another database"},
	}

	for i, t := range targets {
		prefix := fmt.Sprintf("%d. ", i+1)
		if c.target == t.target {
			b.WriteString(c.styles.Focused.Render(prefix + t.name))
		} else {
			b.WriteString(c.styles.Blurred.Render(prefix + t.name))
		}
		b.WriteString(" - " + t.desc + "\n")
	}

	b.WriteString("\n")
	b.WriteString(c.styles.Help.Render("• Press 1-2 to select target • Backspace to go back"))
	return b.String()
}
