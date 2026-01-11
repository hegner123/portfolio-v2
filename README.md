# Portfolio V2

A modern web application built with Go, Templ, and HTMX.

## Tech Stack

- **Go** - Backend server and routing
- **Templ** - Type-safe HTML templating
- **HTMX** - Dynamic interactions without JavaScript

## Getting Started

```bash
# Install dependencies
go mod download

# Install Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Generate templ files
templ generate

# Run the application
go run main.go
```

## Project Structure

```
.
├── ai/              # AI agent documentation and memory
├── handlers/        # HTTP request handlers
├── static/          # Static assets (CSS, JS)
├── templates/       # Templ template files
└── main.go          # Application entry point
```

## Development

See `START.md` for development workflows and common tasks.

## License

MIT
