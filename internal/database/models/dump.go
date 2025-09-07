package models

// DumpConfig holds the configuration for a dump operation
type DumpConfig struct {
	SourceConfig   DatabaseConfig  `json:"source_config"`
	Mode           DumpMode        `json:"mode"`
	Target         DumpTarget      `json:"target"`
	TargetConfig   *DatabaseConfig `json:"target_config,omitempty"` // For direct database imports
	OutputPath     string          `json:"output_path,omitempty"`   // For file exports
	RootTable      string          `json:"root_table,omitempty"`
	RootPrimaryKey string          `json:"root_primary_key,omitempty"`
	IncludeTables  []string        `json:"include_tables,omitempty"`
	ExcludeTables  []string        `json:"exclude_tables,omitempty"`
}

// DumpMode defines the type of dump operation
type DumpMode int

const (
	StructureOnly DumpMode = iota
	StructureAndData
	StructureAndDataExcluding
	StructureAndDataIncludingOnly
)

func (d DumpMode) String() string {
	switch d {
	case StructureOnly:
		return "structure-only"
	case StructureAndData:
		return "structure-and-data"
	case StructureAndDataExcluding:
		return "structure-and-data-excluding"
	case StructureAndDataIncludingOnly:
		return "structure-and-data-including-only"
	default:
		return "unknown"
	}
}

// DumpTarget defines where the output should go
type DumpTarget int

const (
	ToFile DumpTarget = iota
	ToDatabase
)

func (d DumpTarget) String() string {
	switch d {
	case ToFile:
		return "file"
	case ToDatabase:
		return "database"
	default:
		return "unknown"
	}
}
