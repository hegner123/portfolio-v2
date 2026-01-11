.PHONY: build run dev clean templ-generate install-tools tidy status stop

# Build the application
build: templ-generate
	go build -o portfolio-v2 main.go

# Run the application
run: build
	./portfolio-v2

# Development mode - hot reloading with air
dev:
	air

# Generate Go code from Templ files
templ-generate:
	templ generate

# Clean build artifacts
clean:
	rm -f portfolio-v2
	rm -rf tmp
	find . -name "*_templ.go" -type f -delete

# Install development tools
install-tools:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/air-verse/air@latest
	go mod download

# Tidy go.mod
tidy:
	go mod tidy

# Check server status
status:
	@if curl -s http://localhost:8080 > /dev/null 2>&1; then \
		PID=$$(lsof -ti :8080); \
		echo "Server is running (PID: $$PID) on http://localhost:8080"; \
	else \
		echo "Server is not running"; \
	fi

# Stop the server
stop:
	@if curl -s http://localhost:8080 > /dev/null 2>&1; then \
		PID=$$(lsof -ti :8080); \
		echo "Stopping server (PID: $$PID)..."; \
		kill $$PID; \
		echo "Server stopped"; \
	else \
		echo "Server is not running"; \
	fi
