package models

// DatabaseType represents the type of database
type DatabaseType string

const (
	MySQL      DatabaseType = "mysql"
	PostgreSQL DatabaseType = "postgresql"
	SQLite3    DatabaseType = "sqlite3"
)

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Type     DatabaseType `json:"type"`
	Host     string       `json:"host,omitempty"`
	Port     string       `json:"port,omitempty"`
	User     string       `json:"user,omitempty"`
	Password string       `json:"password,omitempty"`
	Database string       `json:"database,omitempty"`
	FilePath string       `json:"file_path,omitempty"` // For SQLite3
}
