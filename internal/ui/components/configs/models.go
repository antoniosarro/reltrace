package configs

import (
	"github.com/antoniosarro/reltrace/internal/database/models"
	"github.com/antoniosarro/reltrace/internal/ui/styles"
	"github.com/charmbracelet/bubbles/textinput"
)

// ConfigCompletedMsg is sent when configuration is completed
type ConfigCompletedMsg struct {
	Config models.DumpConfig
}

// ConfigForm handles database configuration input
type ConfigForm struct {
	styles     *styles.Styles
	inputs     []textinput.Model
	focusIndex int
	dbType     models.DatabaseType
	mode       models.DumpMode
	target     models.DumpTarget
	step       int // 0: db config, 1: mode selection, 2: target selection
}

// NewConfigForm creates a new configuration form
func NewConfigForm(s *styles.Styles) *ConfigForm {
	inputs := make([]textinput.Model, 8)

	for i := range inputs {
		t := textinput.New()
		t.Cursor.Style = s.Cursor
		t.CharLimit = 255

		switch i {
		case 0:
			t.Placeholder = "Database Host (e.g., localhost)"
		case 1:
			t.Placeholder = "Port (e.g., 3306, 5432)"
		case 2:
			t.Placeholder = "Username"
		case 3:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		case 4:
			t.Placeholder = "Database Name"
		case 5:
			t.Placeholder = "Root Table Name"
		case 6:
			t.Placeholder = "Primary Key Value"
		case 7:
			t.Placeholder = "Output Path (optional)"
		}

		inputs[i] = t
	}
	return &ConfigForm{
		styles: s,
		inputs: inputs,
		mode:   models.StructureAndData, // Default mode
		target: models.ToFile,           // Default target
	}
}
