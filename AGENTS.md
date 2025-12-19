# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Development Commands

### Building

```bash
# Build the binary
make build

# Build with Profile-Guided Optimization (PGO)
make build-pgo

# Install to system path
make install

# Install with PGO optimization
make install-pgo

# Generate PGO profile (required before build-pgo)
make profile
```

### Testing

```bash
# Run all tests (excluding large tests)
make test

# Run tests with race detector (quick)
make test-race-short

# Run all tests with race detector
make test-race

# Run large/resource-intensive tests
make test-large

# Run tests with coverage report
make test-coverage

# Run a single test
go test -v -run TestName ./internal/checksum

# Run tests with specific build tags
go test -v -tags=large_tests ./internal/checksum
```

### Code Quality

```bash
# Run linter (golangci-lint)
make lint

# Clean build artifacts
make clean
```

## Architecture Overview

sfvbrr is a high-performance scene release validation tool written in Go. The codebase follows a clean architecture with clear separation of duties.

### Entry Points and Command Structure

1. **Main Entry** (`main.go`) - Minimal entry point that delegates to cmd package
2. **Command Layer** (`cmd/`)
   - `root.go`: Main CLI structure with ASCII banner and global flags
   - `sfv.go`: SFV checksum verification against local folder, or folders
   - `update.go`: Self-update functionality
   - `validate.go`: Various operations performed against folder(s) to check their consistency
   - `version.go`: Version information display

### Core Business Logic

**Checksum Package** (`internal/checksum/`)

- `batch.go`: Parallel batch processing for multiple folders or recursive processing
- `display.go`: User interface and progress display logic
- `progress.go`: Progress tracking for operations
- `sfv.go`: SFV integrity verification, high-performance parallel hashing with adaptive worker pools
- `types.go`: Core data structures

**Preset System** (`internal/preset/`)

- `preset.go`: Configuration preset management
- YAML-based configuration with JSON schema validation
- Supports default settings with preset-specific overrides

**Validation Package** (`internal/validate/`)

- `batch.go`: Parallel batch processing for multiple folders or recursive processing
- `categories.go`: Determining what scene release categories are present
- `display.go`: User interface and progress display logic
- `progress.go`: Progress tracking for operations
- `rules.go`: Validate based on the category and the rules specified in the `presets.yaml`
- `types.go`: Core data structures

### Key Design Patterns and Conventions

1. **Options Pattern**: All operations use structured options for clean APIs
2. **Worker Pools**: Adaptive parallel processing based on workload characteristics
3. **Content Awareness**: Automatic enforcement of content-specific requirements
4. **Error Handling**: User-friendly error messages with actionable feedback
5. **Progress Display**: Real-time feedback with multiple display modes

### Important Implementation Details

1. **Parallel Checksum Validation**:
   - Simple file verification (SFV) is a file format for storing CRC-32 checksums of files to verify the integrity of files
   - Validate the checksum of all files present inside the SFV
   - Adaptive worker count based on file size/count
   - Memory-efficient buffer pooling
   - Optimized for both small and large files

2. **Scene Release Detection**:
   - Detects the type of scene release using the existing library https://github.com/autobrr/rls
   - Use `presets.yaml` to determine what are the rules for each category
   - Based on the type and rules, look at the content of each folder (one or more folders can be specified)
   - Validate the content of the folder(s)

3. **Configuration Files**:
   - Presets: `presets.yaml` for reusable settings
   - Batch: `batch.yaml` for multiple folder operations
   - JSON Schema validation for both formats

4. **File Filtering**:
   - Include/exclude patterns with precedence rules
   - Regex support for complex filtering
   - Default exclusions for common non-media files

5. **Documentation**:
   - Each internal (`checksum`, `preset`, `validate`) will have detailed documentation under the `docs/` folder
   - Main `README.md` in the root folder will have only basic information

### Performance Optimization

- Profile-Guided Optimization (PGO) support via `Makefile`
- Parallel processing with controlled concurrency
- Memory pooling to reduce allocations

### Testing Strategy

- Unit tests alongside implementation files
- Large tests with `large_tests` build tag for expensive operations
- Race condition detection with custom GORACE settings

### Common Development Tasks

```bash
# Add a new release category detection
# Edit internal/validate/categories.go

# Modify category rules
# Edit internal/validate/rules.go

# Add new command
# Create new file in cmd/ following existing patterns

# Test specific functionality
go test -v -run TestValidateRls ./internal/validate
```

### Directory Structure

```tree
sfvbrr/
├── cmd/                  # CLI commands
├── internal/             # Core business logic
│   ├── checksum/         # SFV operations
│   ├── preset/           # Configuration presets
│   └── validate/         # Verification operations
├── schemas/              # JSON schemas for validation
└── test/                 # Test fixtures and data
```

### Git Commit Guidelines

- Use Conventional Commit format: `type(scope): description`
  - Types: `fix`, `feat`, `chore`, `docs`, `test`, `refactor`, `perf`, `style`
  - Example: `fix(checksum): improve efficiency of large file checksum verification`
  - Example: `feat(preset): add support for new release category`
- Keep commits atomic and focused on a single change
- Write clear, descriptive commit messages
- **IMPORTANT**: Never mention Claude or Claude Code in commit messages
