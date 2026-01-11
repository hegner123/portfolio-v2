# Available Actions

This document defines common actions and workflows for the Portfolio V2 project.

## Template Actions

### Create New Component

**When**: Adding a new section or reusable component

**Steps**:
1. Create `.templ` file in `templates/` directory
2. Define component with props if needed
3. Run `templ generate` to create Go code
4. Import and use in parent component or handler

**Example**:
```bash
# Create the template file
# templates/blog-feed.templ

# Generate Go code
templ generate

# Verify generation
ls templates/blog-feed_templ.go
```

**Template Structure**:
```templ
package templates

templ BlogFeed() {
    <section class="blog-feed">
        <h2>Recent Posts</h2>
        <div class="posts-container">
            // Content here
        </div>
    </section>
}
```

### Modify Existing Component

**When**: Updating component markup or structure

**Steps**:
1. Edit the `.templ` file (NOT the `_templ.go` file)
2. Run `templ generate`
3. Test changes in browser

**Command**:
```bash
templ generate && go run main.go
```

### Delete Component

**When**: Removing unused component

**Steps**:
1. Create `trash/` directory if it doesn't exist
2. Move both `.templ` and `_templ.go` files to trash
3. Remove imports from other files
4. Run `templ generate`
5. Test to ensure nothing breaks

**Command**:
```bash
mkdir -p trash
mv templates/unused.templ trash/
mv templates/unused_templ.go trash/
templ generate
```

## Development Actions

### Start Development Server

**When**: Working on the project

**Command**:
```bash
go run main.go
```

**Access**: http://localhost:8080

**Alternative (with auto-reload)**:
```bash
# Install air for live reload (one time)
go install github.com/cosmtrek/air@latest

# Run with air
air
```

### Generate All Templates

**When**: After pulling changes or modifying any `.templ` file

**Command**:
```bash
templ generate
```

**Check for errors**:
```bash
templ generate --log-level=debug
```

### Clean and Rebuild

**When**: Fixing build issues or starting fresh

**Steps**:
```bash
# Clean Go cache
go clean -cache

# Tidy dependencies
go mod tidy

# Regenerate templates
templ generate

# Build
go build -o portfolio-v2 main.go

# Test run
./portfolio-v2
```

## Style Actions

### Create Component Styles

**When**: Adding styles for a new component

**Steps**:
1. Create CSS file: `static/css/component-name.css`
2. Add CSS variables in `:root` if needed
3. Link in `layout.templ`
4. Use BEM naming convention

**Example**:
```bash
# Create file
touch static/css/blog-feed.css

# Add to layout.templ
# <link rel="stylesheet" href="/static/css/blog-feed.css">
```

### Update Existing Styles

**When**: Modifying component appearance

**Steps**:
1. Edit CSS file in `static/css/`
2. Refresh browser (no build needed)
3. Check responsive breakpoints

**Common files**:
- `static/css/base.css` - Global styles
- `static/css/hero.css` - Hero section
- `static/css/about.css` - About section

### Create JavaScript Module

**When**: Adding interactive behavior

**Steps**:
1. Create JS file: `static/js/feature-name.js`
2. Link in `layout.templ` with `defer`
3. Use vanilla JS, avoid frameworks
4. Listen for HTMX events if needed

**Example**:
```bash
# Create file
touch static/js/blog-filter.js

# Add to layout.templ
# <script src="/static/js/blog-filter.js" defer></script>
```

## Handler Actions

### Create New Handler

**When**: Adding new route or API endpoint

**Steps**:
1. Create handler function in `main.go` or `handlers/` directory
2. Register route in `main()`
3. Implement early return pattern
4. Render Templ component or return data

**Example**:
```go
// In main.go or handlers/blog.go
func blogHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    component := templates.BlogFeed()
    if err := component.Render(r.Context(), w); err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        log.Printf("Template rendering error: %v", err)
        return
    }
}

// Register in main()
http.HandleFunc("/blog", blogHandler)
```

### Create HTMX Endpoint

**When**: Adding dynamic content loading

**Steps**:
1. Create handler that returns HTML fragment
2. Use Templ component for response
3. Register route
4. Add HTMX attributes in template

**Example Handler**:
```go
func loadMorePostsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    page := r.URL.Query().Get("page")
    posts := getPostsForPage(page)

    component := templates.PostList(posts)
    component.Render(r.Context(), w)
}
```

**Example Template**:
```html
<button
    hx-get="/api/posts?page=2"
    hx-target="#posts-container"
    hx-swap="beforeend"
>
    Load More
</button>
```

## Content Actions

### Add Blog Post

**When**: Publishing new blog content

**Steps** (once blog system is implemented):
1. Create markdown file in `content/blog/`
2. Add frontmatter (title, date, excerpt)
3. Write content in markdown
4. Restart server to load new content

**File structure**:
```markdown
---
title: "Post Title"
date: 2026-01-11
excerpt: "Brief description"
tags: ["go", "htmx"]
---

Post content here...
```

### Add Project

**When**: Showcasing new project

**Steps** (once project system is implemented):
1. Create markdown file in `content/projects/`
2. Add frontmatter (title, description, tech, links)
3. Add project image to `static/images/projects/`
4. Restart server to load new project

**File structure**:
```markdown
---
title: "Project Name"
description: "What it does"
tech: ["Go", "PostgreSQL", "React"]
github: "https://github.com/user/repo"
demo: "https://project.com"
image: "/static/images/projects/project-name.png"
---

Detailed project description...
```

## Testing Actions

### Manual Browser Testing

**When**: Before committing changes

**Checklist**:
```bash
# 1. Start server
go run main.go

# 2. Test in browsers
# - Chrome: http://localhost:8080
# - Firefox: http://localhost:8080
# - Safari: http://localhost:8080

# 3. Test responsive
# - Use browser dev tools
# - Test mobile viewport (375px)
# - Test tablet viewport (768px)
# - Test desktop viewport (1024px+)

# 4. Check console for errors
# - Open browser console (F12)
# - Look for JavaScript errors
# - Check network tab for failed requests

# 5. Test keyboard navigation
# - Tab through interactive elements
# - Verify focus styles visible
# - Test form submission with Enter

# 6. Test HTMX interactions
# - Verify dynamic content loads
# - Check no full page refreshes
# - Confirm proper error handling
```

### Accessibility Testing

**When**: Adding new interactive features

**Steps**:
```bash
# 1. Keyboard navigation
# - Tab through all interactive elements
# - Verify logical tab order
# - Check focus indicators visible

# 2. Screen reader (macOS VoiceOver)
# - Cmd+F5 to enable
# - Navigate through page
# - Verify labels are descriptive

# 3. Color contrast
# - Use browser DevTools
# - Check text meets WCAG AA standards
# - Verify UI elements distinguishable

# 4. Semantic HTML
# - Inspect DOM structure
# - Verify proper heading hierarchy
# - Check ARIA attributes correct
```

## Git Actions

### Commit Changes

**When**: Completing a feature or fix

**Steps**:
```bash
# 1. Check status
git status

# 2. Review changes
git diff

# 3. Add files (never use git add .)
git add specific-file.go
git add templates/component.templ

# 4. Commit with conventional format
git commit -m "$(cat <<'EOF'
feat(component): add new feature

Detailed description of changes.

Co-Authored-By: Claude Sonnet 4.5 <noreply@anthropic.com>
EOF
)"

# 5. Verify commit
git log -1
```

### Create Feature Branch

**When**: Starting new feature work

**Steps**:
```bash
# Create and switch to new branch
git checkout -b feature/blog-feed

# Work on feature...

# Push to remote
git push -u origin feature/blog-feed
```

### Update Implementation Status

**When**: Completing a phase or task

**Steps**:
1. Open `ai/IMPLEMENTATION.md`
2. Update task checkboxes `[ ]` to `[x]`
3. Update phase status
4. Update "Last Updated" date
5. Commit changes

**Example**:
```bash
# Edit IMPLEMENTATION.md
# Change status from "Not Started" to "In Progress" or "Completed"
# Mark completed tasks with [x]

git add ai/IMPLEMENTATION.md
git commit -m "docs(implementation): update Phase 5 status"
```

## Build Actions

### Development Build

**When**: Testing locally

**Command**:
```bash
templ generate && go run main.go
```

### Production Build

**When**: Preparing for deployment

**Commands**:
```bash
# Generate templates
templ generate

# Tidy dependencies
go mod tidy

# Build with optimizations
go build -ldflags="-s -w" -o portfolio-v2 main.go

# Test production build
./portfolio-v2
```

**Flags explained**:
- `-s`: Strip symbol table
- `-w`: Strip DWARF debugging info
- `-o`: Output filename

### Docker Build (Future)

**When**: Containerizing application

**Dockerfile pattern**:
```dockerfile
FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -ldflags="-s -w" -o portfolio-v2 main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/portfolio-v2 .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./portfolio-v2"]
```

## Maintenance Actions

### Update Dependencies

**When**: Monthly or when security updates available

**Commands**:
```bash
# Check for updates
go list -u -m all

# Update specific dependency
go get github.com/a-h/templ@latest

# Update all dependencies
go get -u ./...

# Tidy up
go mod tidy

# Test everything still works
templ generate
go run main.go
```

### Clean Project

**When**: Build issues or disk space cleanup

**Commands**:
```bash
# Remove generated files
mkdir -p trash
mv templates/*_templ.go trash/

# Clean Go cache
go clean -cache -modcache -testcache

# Regenerate everything
templ generate
go mod download
go build -o portfolio-v2 main.go
```

### Update Templ CLI

**When**: New Templ version available

**Commands**:
```bash
# Update Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Verify version
templ version

# Regenerate templates with new version
templ generate
```

## Deployment Actions

### Deploy to Production

**When**: Ready to deploy new version

**Pre-deployment checklist**:
- [ ] All tests passing
- [ ] Templates generated (`templ generate`)
- [ ] Dependencies tidy (`go mod tidy`)
- [ ] Production build successful
- [ ] Environment variables configured
- [ ] Static assets included
- [ ] Database migrations applied (if applicable)
- [ ] Backup current version

**Deployment steps** (platform-specific):
```bash
# 1. Build production binary
templ generate
go build -ldflags="-s -w" -o portfolio-v2 main.go

# 2. Upload to server
scp portfolio-v2 user@server:/app/
scp -r static user@server:/app/

# 3. Restart service
ssh user@server "systemctl restart portfolio"

# 4. Verify deployment
curl https://yoursite.com
```

### Rollback Deployment

**When**: Production issues detected

**Steps**:
```bash
# 1. SSH to server
ssh user@server

# 2. Stop current version
systemctl stop portfolio

# 3. Restore previous version
cp /app/backups/portfolio-v2.backup /app/portfolio-v2

# 4. Restart service
systemctl start portfolio

# 5. Verify rollback
curl http://localhost:8080
```

## Debug Actions

### Debug Template Rendering

**When**: Template not rendering correctly

**Steps**:
```bash
# 1. Verify template generated
ls -la templates/*_templ.go

# 2. Check for syntax errors
templ generate --log-level=debug

# 3. Check handler logs
go run main.go
# Look for "Template rendering error" in output

# 4. Verify component import
grep "templates\." main.go
```

### Debug HTMX Issues

**When**: HTMX not working as expected

**Steps**:
```bash
# 1. Enable HTMX logging in browser console
# Add to layout.templ:
# <script>htmx.logAll();</script>

# 2. Check network tab for HTMX requests
# - Verify request URL correct
# - Check response is HTML fragment
# - Verify response status 200

# 3. Check HTMX attributes
# - Verify hx-target exists
# - Check hx-swap value correct
# - Confirm hx-trigger appropriate
```

### Debug CSS Not Loading

**When**: Styles not applying

**Steps**:
```bash
# 1. Verify file exists
ls static/css/component.css

# 2. Check link in layout.templ
grep "component.css" templates/layout.templ

# 3. Check browser network tab
# - Verify CSS file requested
# - Check response status 200
# - Verify correct Content-Type

# 4. Check static file server
# - Verify /static/ route registered
# - Check file server directory correct
```

## Performance Actions

### Optimize Images

**When**: Adding new images

**Commands**:
```bash
# Install ImageMagick (one time)
brew install imagemagick

# Optimize JPEG
convert input.jpg -quality 85 -strip output.jpg

# Optimize PNG
pngquant input.png --output output.png --quality 85-95

# Convert to WebP
cwebp -q 85 input.jpg -o output.webp
```

### Analyze Bundle Size

**When**: Concerned about page weight

**Steps**:
```bash
# Check static file sizes
du -sh static/css/*
du -sh static/js/*
du -sh static/images/*

# Check total static directory
du -sh static/

# Identify large files
find static -type f -size +100k -exec ls -lh {} \;
```

### Profile Server Performance

**When**: Investigating slow requests

**Steps**:
```go
// Add to main.go
import _ "net/http/pprof"

// Access profiler
// http://localhost:8080/debug/pprof/
```

## Documentation Actions

### Update Implementation Plan

**When**: Completing tasks or starting new phase

**File**: `ai/IMPLEMENTATION.md`

**Steps**:
1. Mark completed tasks `[x]`
2. Update phase status
3. Update "Current Phase" and "Next Action"
4. Update "Last Updated" date
5. Add implementation details if significant
6. Commit changes

### Document New Convention

**When**: Establishing new pattern or standard

**File**: `ai/conventions.md`

**Steps**:
1. Identify relevant section
2. Add new convention with example
3. Explain rationale if not obvious
4. Commit changes

### Update Overview

**When**: Major architectural changes or new features

**File**: `ai/overview.md`

**Steps**:
1. Update relevant sections
2. Add new features to "Current Features"
3. Update directory structure if changed
4. Commit changes
