# Reltrace

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)

**Reltrace** is a powerful command-line database export tool that maintains referential integrity by following foreign key relationships. It supports selective data extraction, complete database dumps, and direct database-to-database transfers without external dependencies.

---

## Features

- **Multi-Database Support** – Works with MySQL, PostgreSQL, and SQLite3
- **Pure Go Implementation** – No external tools required (no mysqldump, pg_dump dependencies)
- **Referential Integrity** – Automatically follows foreign key relationships
- **Flexible Export Modes** – Structure only, full database, selective inclusion/exclusion
- **Direct Database Transfer** – Export directly to another database without intermediate files
- **Circular Reference Handling** – Safely handles complex relationship cycles
- **Interactive TUI** – User-friendly terminal interface for configuration
- **Batch Processing** – Optimized for large datasets with efficient memory usage

---

## Export Modes

1. **Structure Only** – Export database schema without data
2. **Structure + All Data** – Complete database backup
3. **Structure + Data (Excluding)** – Export everything except records related to a specific root
4. **Structure + Data (Including Only)** – Export only records related to a specific root

## Export Targets

- **File Export** – Generate SQL dump files
- **Database Export** – Direct transfer to another database

---

## Installation

### Using Go Install
```bash
go install github.com/antoniosarro/reltrace/cmd/reltrace@latest
```

### Building from Source
```bash
git clone https://github.com/antoniosarro/reltrace.git
cd reltrace
make build
```

### Using Nix (with flakes)
```bash
nix develop
make build
```

## Usage
### Interactive mode
```bash
./bin/reltrace
```

The interactive TUI will guide you through:
1. Database type selection
2. Connection configuration
3. Export mode selection
4. Target configuration

## Example Use Cases
### Complete Database Backup:
- Export entire database structure and data to SQL file

### Selective Data Extraction:
- Extract a customer record and all related orders, payments, and history
- Export a project with all associated tasks, comments, and files

### Database Migration:
- Transfer specific datasets between environments
- Copy production subsets to development databases

### Data cleanup:
- Export everything except test data or specific user records

## Supported Databases
| **Database** | **Structure Export** | **Data Export** | **Direct Transfer** |
|:------------:|:--------------------:|:---------------:|:-------------------:|
|     MySQL    |           ✅          |        ✅        |          ✅          |
|  PostgreSQL  |           ✅          |        ✅        |          ✅          |
|    SQLite3   |           ✅          |        ✅        |          ✅          |

## Tech Stack
- **Language**: Go 1.24+
- **Database Drivers**:
    - MySQL: `github.com/go-sql-driver/mysql`
    - PostgreSQL: `github.com/lib/pq`
    - SQLite3: `github.com/mattn/go-sqlite3`
- **TUI Framework**: Bubble Tea
- **Build System**: Make + Nix Flakes
- **License**: MIT

## Development
### Prerequisites
- Go 1.24 or later
- Make
- Database servers for testing

### Setup Development Environment
```bash
# Clone repository
git clone https://github.com/antoniosarro/reltrace.git
cd reltrace

# Install dependencies
make dev-deps

# Run tests
make test

# Build
make build

# Run
make dev
```

### Using Nix
```bash
# Enter development shell
nix develop

# All development tools are now available
make build
```

### Project Struture
```
reltrace/
├── cmd/reltrace/           # Main application entry point
├── internal/
│   ├── app/reltrace/       # Application setup
│   ├── config/             # Configuration management
│   ├── database/
│   │   ├── adapters/       # Database-specific implementations
│   │   ├── engine/         # Core dump engine
│   │   ├── models/         # Data structures
│   │   └── processor/      # Legacy compatibility layer
│   └── ui/                 # Terminal user interface
│       ├── components/     # UI components
│       └── styles/         # Styling definitions
├── db.sql                  # Test database schema
└── Makefile               # Build automation
```

## Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Authors
[![][contributors]][contributors-graph]

<!----------------------------------{ Labels }--------------------------------->
[contributors]: https://contrib.rocks/image?repo=antoniosarro/reltrace
[contributors-graph]: https://github.com/antoniosarro/reltrace/graphs/contributors