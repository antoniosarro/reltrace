package config

// AppConfig holds the application configuration
type AppConfig struct {
	Output OutputConfig
	UI     UIConfig
}

// OutputConfig defines output settings
type OutputConfig struct {
	Directory string
	Format    string
	Timestamp bool
}

// UIConfig defines UI settings
type UIConfig struct {
	Theme           string
	AnimateProgress bool
}

// DefaultConfig returns a default configuration
func DefaultConfig() AppConfig {
	return AppConfig{
		Output: OutputConfig{
			Directory: ".",
			Format:    "sql",
			Timestamp: true,
		},
		UI: UIConfig{
			Theme:           "default",
			AnimateProgress: true,
		},
	}
}
