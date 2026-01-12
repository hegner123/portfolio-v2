# Portfolio V2

A modern portfolio website with blog and project management CMS, built with Go, Templ, and HTMX.

**Live:** [michaelhegner.com](https://michaelhegner.com)

## Built in a Day

This entire production-ready application—including a blog and project CMS with admin dashboard—was built and deployed in **~10 hours of active development** (January 11, 2026), working with Claude Code.

**Timeline:**
- **8:39 AM** - Started with interactive hero section (momentum-based grid physics animation)
- **10:30 AM** - About section with 50+ technologies and expandable skills grid
- **3:50 PM** - Navigation system with smooth scrolling and mobile menu
- **4:36 PM** - Blog feed with tag filtering, project showcase with filtering, contact section
- **4:51 PM** - Admin dashboard with authentication, CRUD operations for blog posts and projects
- **10:38 PM** - Complete deployment pipeline with rsync, cross-compilation, and HTTPS

**What was delivered:**
- ✅ Full-stack Go application with type-safe templating
- ✅ Interactive hero section with custom momentum-based physics animation
- ✅ Blog and project CMS with admin dashboard (create, edit, delete posts and projects)
- ✅ HTTP Basic Authentication for admin routes
- ✅ Blog system with SQLite database, markdown support, tag filtering, and HTMX pagination
- ✅ Project showcase with category filtering, dynamic HTMX loading, and featured project badges
- ✅ Responsive design with animated mobile navigation and smooth scrolling
- ✅ Expandable skills grid with JavaScript animations
- ✅ Complete deployment infrastructure (rsync, systemd, Caddy)
- ✅ Production deployment on Linode with automatic HTTPS
- ✅ Security headers, SEO optimization, and accessibility features

This demonstrates rapid, production-quality development using AI-assisted workflows while maintaining best practices, proper architecture, and comprehensive documentation.

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
