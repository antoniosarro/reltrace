package database

import (
	"github.com/antoniosarro/reltrace/internal/database/models"
	"github.com/antoniosarro/reltrace/internal/ui/styles"
)

// DatabaseSelectedMsg is sent when a database type is selected
type DatabaseSelectedMsg struct {
	DatabaseType models.DatabaseType
}

// DatabaseSelector handles database type selection
type DatabaseSelector struct {
	styles   *styles.Styles
	choices  []models.DatabaseType
	selected int
}

// NewDatabaseSelector creates a new database selector
func NewDatabaseSelector(s *styles.Styles) *DatabaseSelector {
	return &DatabaseSelector{
		styles: s,
		choices: []models.DatabaseType{
			models.MySQL,
			models.PostgreSQL,
			models.SQLite3,
		},
		selected: 0,
	}
}
