package ui

import (
	"github.com/antoniosarro/reltrace/internal/config"
	"github.com/antoniosarro/reltrace/internal/ui/components/configs"
	"github.com/antoniosarro/reltrace/internal/ui/components/database"
	"github.com/antoniosarro/reltrace/internal/ui/styles"
)

// ViewState represents the current view state
type ViewState int

const (
	DatabaseSelectionView ViewState = iota
	ConfigurationView
	ProcessingView
	CompletionView
)

// Model represents the main UI model
type Model struct {
	state  ViewState
	config config.AppConfig
	styles *styles.Styles

	// Components
	configForm *configs.ConfigForm
	dbSelector *database.DatabaseSelector

	// State
	error string

	// Window size for proper rendering
	width  int
	height int
}

// New creates a new UI model
func New(config config.AppConfig) *Model {
	s := styles.New()

	return &Model{
		state:      DatabaseSelectionView,
		config:     config,
		styles:     s,
		configForm: configs.NewConfigForm(s),
		dbSelector: database.NewDatabaseSelector(s),
	}
}
