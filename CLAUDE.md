# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Portfolio V2 is a single-page portfolio website built with Go, Templ (type-safe HTML templating), and HTMX for dynamic interactions. The project is in early development with the basic server structure in place.

## Tech Stack

- **Go 1.25.5** - Backend HTTP server
- **Templ** - Type-safe HTML templating that generates Go code
- **HTMX** - Client-side dynamic interactions without JavaScript frameworks

## Essential Commands

### Setup and Dependencies
```bash
# Install dependencies
go mod download

# Install Templ CLI (required for development)
go install github.com/a-h/templ/cmd/templ@latest

# Generate Go code from .templ files (run after modifying templates)
templ generate
```

### Development
```bash
# Run the application
go run main.go
# Server runs on http://localhost:8080

# Clean up go.mod
go mod tidy
```

### Building
```bash
# Build the application
go build -o portfolio-v2 main.go
```

## Architecture

### Current Structure
The application is a simple HTTP server (main.go:9-23) with:
- Static file server mounted at `/static/` (main.go:11-12)
- Route handlers registered with `http.HandleFunc`
- Currently only has a home handler (main.go:15) serving temporary HTML

### Template System (Templ)
- `.templ` files contain type-safe Go templates
- **CRITICAL**: After creating or modifying `.templ` files, run `templ generate` to generate the corresponding `_templ.go` files
- The generated Go code is what gets imported and used in handlers
- Templ files should live in the `templates/` directory

### Handler Pattern
- Handlers should be organized in the `handlers/` directory (currently empty)
- Current pattern: Early return for method validation (main.go:26-29)
- Handlers will eventually render Templ components instead of raw HTML

### Static Assets
- Served from `./static/` directory
- Accessible at `/static/*` URLs
- Directory exists but is currently empty

## Project Roadmap

The implementation plan is tracked in `/Users/home/Documents/Michael/portfolio-v2/ai/IMPLEMENTATION.md`. Key sections to be built:
1. Core application structure (in progress)
2. Hero section with animation
3. About section
4. Blog feed (with HTMX dynamic loading)
5. Project feed (with HTMX dynamic loading)
6. Contact form (with HTMX submission)
7. Navigation with smooth scrolling

## Known Issues

- `github.com/a-h/templ` is listed in go.mod but not yet used (needs `go mod tidy` or actual usage)
- Home handler currently returns hardcoded HTML instead of using Templ templates

## Development Workflow

1. Create `.templ` files in `templates/` directory
2. Run `templ generate` to create Go code
3. Import and use generated components in handlers
4. Test by running `go run main.go` and visiting http://localhost:8080
