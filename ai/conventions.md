# Coding Conventions

## Go Conventions

### General Style

- Follow standard Go formatting (`gofmt`)
- Use `any` instead of `interface{}` in type definitions
- Prefer early return pattern over nested conditionals
- Avoid deep nesting - keep code flat
- Use absolute paths or `$HOME` instead of `~` shorthand

### Package Structure

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "portfolio-v2/templates"  // Local imports last
)
```

### Handler Pattern

```go
func exampleHandler(w http.ResponseWriter, r *http.Request) {
    // Early returns for validation
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Process request
    component := templates.Example()

    // Render with error handling
    if err := component.Render(r.Context(), w); err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        log.Printf("Template rendering error: %v", err)
        return
    }
}
```

### Error Handling

- Always check errors
- Log errors with context
- Return appropriate HTTP status codes
- Use `log.Printf` for server-side logging

```go
if err := doSomething(); err != nil {
    log.Printf("Failed to do something: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
}
```

### Deprecated Packages

- Do NOT use `io/ioutil` (deprecated)
- Use `io` and `os` packages instead

### Integer Range Loops (Go 1.22+)

```go
// Use modern range syntax
for i := range 10 {
    // i goes from 0 to 9
}
```

## Templ Conventions

### File Naming

- Component files: `componentname.templ` (lowercase)
- Generated files: `componentname_templ.go` (never edit these)

### Component Structure

```templ
package templates

templ ComponentName() {
    <div class="component-name">
        <h2>Component Title</h2>
        <div class="component-content">
            { children... }
        </div>
    </div>
}
```

### Component with Props

```templ
package templates

type ComponentProps struct {
    Title   string
    Content string
}

templ ComponentWithProps(props ComponentProps) {
    <div class="component">
        <h2>{ props.Title }</h2>
        <p>{ props.Content }</p>
    </div>
}
```

### Component Composition

```templ
templ Layout(title string) {
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <title>{ title }</title>
        </head>
        <body>
            { children... }
        </body>
    </html>
}

templ Home() {
    @Layout("Portfolio") {
        @Hero()
        @About()
    }
}
```

### Best Practices

- Keep components focused and single-purpose
- Use props for dynamic content
- Compose larger components from smaller ones
- Always run `templ generate` after changes
- Don't edit `*_templ.go` files manually

## CSS Conventions

### File Organization

- One CSS file per component: `component-name.css`
- Global styles in `main.css` or `base.css`
- Place in `/static/css/` directory

### Naming Convention

Use BEM-like naming:

```css
/* Component */
.component-name { }

/* Element */
.component-name__element { }

/* Modifier */
.component-name--modifier { }

/* State */
.component-name.is-active { }
```

### CSS Variables for Theming

```css
:root {
    /* Colors */
    --color-primary: #8b5cf6;
    --color-secondary: #3b82f6;
    --color-background: #0a0a0a;
    --color-text: #e5e5e5;

    /* Spacing */
    --spacing-xs: 0.5rem;
    --spacing-sm: 1rem;
    --spacing-md: 1.5rem;
    --spacing-lg: 2rem;
    --spacing-xl: 3rem;

    /* Typography */
    --font-family-base: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    --font-size-base: 1rem;
    --font-size-lg: 1.25rem;
    --font-size-xl: 1.5rem;

    /* Effects */
    --shadow-sm: 0 2px 4px rgba(0, 0, 0, 0.1);
    --shadow-md: 0 4px 8px rgba(0, 0, 0, 0.2);
    --shadow-glow: 0 0 20px rgba(139, 92, 246, 0.3);
}
```

### Responsive Design

```css
/* Mobile-first approach */
.component {
    /* Base styles for mobile */
}

/* Tablet and up */
@media (min-width: 768px) {
    .component {
        /* Tablet styles */
    }
}

/* Desktop and up */
@media (min-width: 1024px) {
    .component {
        /* Desktop styles */
    }
}
```

### Animation Guidelines

```css
/* Use transitions for simple state changes */
.button {
    transition: background-color 0.3s ease;
}

/* Use animations for complex sequences */
@keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
}

.element {
    animation: fadeIn 0.5s ease-in-out;
}
```

## JavaScript Conventions

### File Organization

- One JS file per feature: `feature-name.js`
- Place in `/static/js/` directory
- Load with `defer` attribute

### Code Style

```javascript
// Use const/let, never var
const elements = document.querySelectorAll('.element');
let count = 0;

// Use arrow functions for callbacks
elements.forEach(el => {
    el.addEventListener('click', handleClick);
});

// Early returns
function handleClick(event) {
    if (!event.target) return;

    // Process event
}
```

### DOM Manipulation

```javascript
// Cache DOM queries
const hero = document.querySelector('.hero');
const squares = hero.querySelectorAll('.square');

// Use classList for class manipulation
element.classList.add('is-active');
element.classList.remove('is-hidden');
element.classList.toggle('is-expanded');
```

### Event Listeners

```javascript
// Named functions for listeners (easier to debug)
function handleMouseMove(event) {
    // Handle event
}

element.addEventListener('mousemove', handleMouseMove);

// Cleanup when needed
element.removeEventListener('mousemove', handleMouseMove);
```

### HTMX Integration

```javascript
// Listen to HTMX events
document.body.addEventListener('htmx:afterSwap', (event) => {
    // Reinitialize components after HTMX swap
    initializeNewContent(event.detail.target);
});
```

## HTMX Conventions

### Attribute Naming

```html
<!-- Use semantic HTMX attributes -->
<button
    hx-get="/api/posts"
    hx-target="#posts-container"
    hx-swap="innerHTML"
    hx-trigger="click"
>
    Load Posts
</button>
```

### Common Patterns

#### Load More

```html
<button
    hx-get="/api/posts?page=2"
    hx-target="#posts-container"
    hx-swap="beforeend"
>
    Load More
</button>
```

#### Form Submission

```html
<form
    hx-post="/contact"
    hx-target="#form-response"
    hx-swap="innerHTML"
>
    <input name="email" type="email" required>
    <button type="submit">Submit</button>
</form>
```

#### Infinite Scroll

```html
<div
    hx-get="/api/posts?page=2"
    hx-trigger="revealed"
    hx-swap="afterend"
>
    Loading...
</div>
```

### Response Handling

Server should return HTML fragments:

```go
func postsHandler(w http.ResponseWriter, r *http.Request) {
    component := templates.PostList(posts)
    component.Render(r.Context(), w)
}
```

## File Naming Conventions

### General Rules

- Use lowercase with hyphens: `file-name.ext`
- Be descriptive: `hero-animation.js` not `anim.js`
- Group by feature/component

### Examples

```
templates/
├── layout.templ
├── home.templ
├── hero.templ
├── about.templ
├── blog-feed.templ
└── contact-form.templ

static/css/
├── base.css
├── hero.css
├── about.css
├── blog-feed.css
└── contact-form.css

static/js/
├── hero-animation.js
├── about-skills.js
└── htmx-helpers.js
```

## Git Conventions

### Commit Messages

Use conventional commit format:

```
type(scope): brief description

Detailed explanation if needed.
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting, CSS changes
- `refactor`: Code restructuring
- `test`: Adding tests
- `chore`: Maintenance

Examples:
```
feat(hero): add momentum-based grid animation
fix(about): expand/collapse animation glitch
docs(ai): create conventions documentation
style(about): adjust skill card glow effect
```

### Co-authoring with Claude

All commits include:
```
Co-Authored-By: Claude Sonnet 4.5 <noreply@anthropic.com>
```

### File Management

- **Never delete files** - move to trash directory
- Check file existence before copying
- Never use regex for file editing/deletion operations

## Security Conventions

### Input Validation

```go
// Always validate user input
func validateEmail(email string) bool {
    // Validation logic
    return isValid
}

// Sanitize before processing
cleaned := sanitizeInput(userInput)
```

### XSS Prevention

- Templ auto-escapes content (safe by default)
- Never use raw HTML injection
- Validate all user inputs server-side

### Error Messages

```go
// Don't expose internal errors to users
if err != nil {
    log.Printf("Database error: %v", err)  // Log details
    http.Error(w, "An error occurred", http.StatusInternalServerError)  // Generic message
    return
}
```

## Accessibility Conventions

### Semantic HTML

```html
<!-- Use proper semantic elements -->
<nav>
    <ul>
        <li><a href="#about">About</a></li>
    </ul>
</nav>

<main>
    <section id="about">
        <h2>About Me</h2>
    </section>
</main>
```

### ARIA Labels

```html
<!-- Decorative elements -->
<div class="decoration" aria-hidden="true"></div>

<!-- Interactive elements -->
<button aria-label="Expand skills" aria-expanded="false">
    Show More
</button>

<!-- Landmarks -->
<section aria-label="Professional Skills">
    <!-- Content -->
</section>
```

### Keyboard Navigation

- All interactive elements must be keyboard accessible
- Use `tabindex` appropriately
- Provide visible focus states

```css
.button:focus-visible {
    outline: 2px solid var(--color-primary);
    outline-offset: 2px;
}
```

## Documentation Conventions

### Code Comments

Use comments sparingly - prefer self-documenting code:

```go
// GOOD - explain why, not what
// Exponential scaling creates more dramatic parting effect
force *= partAmount * partAmount

// BAD - obvious from code
// Multiply force by partAmount squared
force *= partAmount * partAmount
```

### File Headers

Minimal headers, focus on purpose:

```go
// Package templates contains Templ components for the portfolio site.
package templates
```

### TODO Comments

```go
// TODO(phase-6): Integrate with GitHub API for project data
// FIXME: Animation jank on mobile devices
// NOTE: This must run after HTMX swap
```

## Testing Conventions

### Manual Testing Checklist

Before committing:
- [ ] Test on Chrome, Firefox, Safari
- [ ] Test on mobile viewport
- [ ] Test keyboard navigation
- [ ] Test with screen reader (if applicable)
- [ ] Check console for errors
- [ ] Verify HTMX interactions

### Future Automated Testing

Structure for when tests are added:

```
tests/
├── handlers/
│   └── home_test.go
├── templates/
│   └── components_test.go
└── integration/
    └── flow_test.go
```

## Build Conventions

### Development Build

```bash
# Always generate templates first
templ generate

# Run with go run for development
go run main.go
```

### Production Build

```bash
# Generate templates
templ generate

# Build with optimizations
go build -ldflags="-s -w" -o portfolio-v2 main.go
```

### Deployment Checklist

- [ ] Run `templ generate`
- [ ] Run `go mod tidy`
- [ ] Test build locally
- [ ] Check static assets are included
- [ ] Verify environment variables
- [ ] Test production build

## Editor Configuration

### Recommended VS Code Extensions

- Go (official Go extension)
- Templ (templ-go.templ-vscode)
- HTMX Attributes (otovo-oss.htmx-tags)

### Settings

```json
{
    "go.formatTool": "gofmt",
    "editor.formatOnSave": true,
    "[templ]": {
        "editor.defaultFormatter": "a-h.templ"
    }
}
```
