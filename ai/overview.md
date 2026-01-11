# Project Overview

## Description

Portfolio V2 is a modern single-page portfolio website showcasing professional work, blog posts, and projects. Built with Go, Templ, and HTMX, it emphasizes server-side rendering with dynamic client-side interactions without heavy JavaScript frameworks.

## Goals

- **Performance**: Fast page loads with minimal JavaScript
- **Maintainability**: Type-safe templates with Go
- **Interactivity**: Dynamic content loading without full page refreshes
- **Accessibility**: Semantic HTML with proper ARIA labels
- **Responsiveness**: Mobile-first design approach

## Tech Stack

### Backend
- **Go 1.25.5**: HTTP server and application logic
- **net/http**: Standard library HTTP server
- **Templ**: Type-safe HTML templating that compiles to Go code

### Frontend
- **HTMX**: Dynamic HTML updates without writing JavaScript
- **Vanilla CSS**: Custom CSS with CSS variables for theming
- **Vanilla JavaScript**: Minimal JS for animations and interactions

### Development Tools
- **Templ CLI**: Template compilation (`templ generate`)
- **Go modules**: Dependency management

## Architecture

### Request Flow
1. Client requests `/` route
2. `homeHandler` in `main.go` processes request
3. Handler renders Templ components (Layout → Home → Hero, About, etc.)
4. Server returns complete HTML page
5. HTMX intercepts future interactions for dynamic updates

### Component Structure
```
Layout (base HTML structure)
└── Home (main page container)
    ├── Hero (animated grid section)
    ├── About (bio and skills)
    ├── BlogFeed (dynamic list - upcoming)
    ├── ProjectFeed (dynamic grid - upcoming)
    └── Contact (form - upcoming)
```

### Static Assets
- **CSS**: `/static/css/` - Styles for each section
- **JS**: `/static/js/` - Interactive scripts (animations, HTMX helpers)
- **Images**: `/static/images/` - Static images and assets

## Directory Structure

```
portfolio-v2/
├── ai/                     # AI documentation and planning
│   ├── IMPLEMENTATION.md   # Implementation tracker
│   ├── overview.md         # This file
│   └── conventions.md      # Coding conventions
├── handlers/               # HTTP handlers (future organization)
├── templates/              # Templ component files
│   ├── layout.templ        # Base HTML structure
│   ├── home.templ          # Main page container
│   ├── hero.templ          # Hero section
│   ├── about.templ         # About section
│   └── *_templ.go          # Generated Go files (do not edit)
├── static/                 # Static assets
│   ├── css/                # Stylesheets
│   ├── js/                 # JavaScript files
│   └── images/             # Images and graphics
├── main.go                 # Application entry point
├── go.mod                  # Go dependencies
├── CLAUDE.md               # Claude Code instructions
└── START.md                # Onboarding guide
```

## Current Features

### Completed (Phases 1-6, 8)

#### Phase 1-2: Foundation
- Go HTTP server on port 8080
- Static file serving at `/static/`
- Templ integration with proper generation
- Base layout template with HTML structure

#### Phase 3: Hero Section
- **Interactive Grid Animation**
  - Semi-transparent squares arranged in grid
  - Attraction mode: squares gravitate to mouse
  - Parting mode: hold mouse to make squares flee
  - Momentum-based physics with damping
  - Easter egg: 10 rapid clicks = explosion effect
  - Dynamic opacity based on interaction
- Name and tagline display
- Fully responsive
- Accessible (aria-hidden decorative elements)

#### Phase 4: About Section
- **Professional Bio**: Personalized content about approach
- **Skills Grid**: 9 categories with 50+ technologies
  - Backend (Go, Node.js, Python, etc.)
  - Frontend (React, Next.js, HTMX, etc.)
  - Databases (PostgreSQL, MySQL, MongoDB, etc.)
  - AI & LLM Integration
  - Cloud & Infrastructure
  - CMS Platforms
  - Tools & Platforms
  - Security & Compliance
  - Practices & Methodologies
- **Expand/Collapse**: Show top 3 skills, expand for all
- **Dark Theme**: Purple/blue gradient with glow effects
- **Fully Responsive**: Single column on mobile
- **Interactive**: Smooth JavaScript transitions

#### Phase 8: Navigation & Smooth Scrolling
- **Sticky Navigation**: Glassmorphic nav bar with backdrop blur
- **Smooth Scrolling**: CSS-based with scroll-padding for sticky nav offset
- **Active Section Highlighting**: Intersection Observer API detection
- **Mobile Navigation**: Hamburger menu with slide-in panel (≤768px)
- **Styling**: Purple/blue gradient accents, hover lift effects, scrolled state with glow
- **Accessibility**: Keyboard navigation, ARIA labels, ESC key support, focus indicators
- **Performance**: RequestAnimationFrame for scroll updates, passive event listeners
- **Features**:
  - 5 navigation links (Home, About, Blog, Projects, Contact)
  - Automatic active section detection and highlighting
  - Mobile menu toggle with backdrop overlay
  - Scrolled state detection for enhanced background
  - Responsive design with touch-optimized mobile menu

### Upcoming (Phases 7, 9-11)

- **Phase 7**: Contact Form with HTMX submission
- **Phase 9**: Design polish and accessibility
- **Phase 10**: Performance optimization
- **Phase 11**: Deployment

## Development Workflow

### Making Changes

1. **Modify Templates**
   ```bash
   # Edit .templ files
   vim templates/about.templ

   # Generate Go code
   templ generate
   ```

2. **Update Styles**
   ```bash
   # Edit CSS files directly
   vim static/css/about.css
   ```

3. **Run Server**
   ```bash
   go run main.go
   # Visit http://localhost:8080
   ```

4. **Build for Production**
   ```bash
   go build -o portfolio-v2 main.go
   ```

### Key Commands

```bash
# Install dependencies
go mod download

# Install Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Generate templates
templ generate

# Run development server
go run main.go

# Clean dependencies
go mod tidy

# Build binary
go build -o portfolio-v2 main.go
```

## Design Philosophy

### Simplicity Over Complexity
- Use standard library where possible
- Avoid unnecessary abstractions
- Keep components focused and single-purpose

### Server-First
- Render HTML on server
- Use HTMX for dynamic updates
- Minimize client-side JavaScript

### Progressive Enhancement
- Works without JavaScript (basic functionality)
- Enhanced with HTMX and JS for better UX
- Accessible to all users

### Type Safety
- Templ provides compile-time template checking
- Go's type system prevents runtime errors
- Clear interfaces between components

## Performance Considerations

- **Single-page app**: All sections load initially (fast subsequent navigation)
- **Lazy loading**: Images and heavy content load on demand
- **Minimal JavaScript**: Only what's needed for interactions
- **CSS optimization**: Use CSS variables and efficient selectors
- **HTMX caching**: Reuse HTML fragments when possible

## Security Considerations

- **Input validation**: All form inputs validated server-side
- **XSS prevention**: Templ auto-escapes content
- **CSRF protection**: To be implemented for forms
- **Rate limiting**: To be implemented for contact form
- **Secure headers**: To be configured in production

## Future Enhancements

- Blog post markdown rendering
- Project data from GitHub API
- Email integration for contact form
- Dark/light theme toggle
- Analytics integration
- RSS feed for blog
- Search functionality
