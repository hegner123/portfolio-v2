package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and creates tables
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("create tables: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS blog_posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		excerpt TEXT NOT NULL,
		content TEXT NOT NULL,
		published_at DATETIME NOT NULL,
		tags TEXT NOT NULL DEFAULT '[]',
		author TEXT NOT NULL DEFAULT 'Michael',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_blog_posts_slug ON blog_posts(slug);
	CREATE INDEX IF NOT EXISTS idx_blog_posts_published_at ON blog_posts(published_at DESC);

	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		description TEXT NOT NULL,
		technologies TEXT NOT NULL DEFAULT '[]',
		github_url TEXT NOT NULL,
		image_url TEXT NOT NULL,
		featured INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_projects_slug ON projects(slug);
	CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects(featured DESC, created_at DESC);

	CREATE TABLE IF NOT EXISTS contact_submissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		message TEXT NOT NULL,
		submitted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		ip_address TEXT,
		user_agent TEXT
	);

	CREATE INDEX IF NOT EXISTS idx_contact_submissions_submitted_at ON contact_submissions(submitted_at DESC);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("execute schema: %w", err)
	}

	return nil
}

// SeedBlogPosts adds sample blog posts for development
func SeedBlogPosts(db *sql.DB) error {
	count, err := CountBlogPosts(db, "")
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Blog posts already exist, skipping seed")
		return nil
	}

	posts := []struct {
		title       string
		slug        string
		excerpt     string
		content     string
		publishedAt time.Time
		tags        []string
	}{
		{
			title:   "Building with Go and HTMX",
			slug:    "building-with-go-and-htmx",
			excerpt: "Exploring how Go and HTMX create a powerful combination for modern web development without heavy JavaScript frameworks.",
			content: `# Building with Go and HTMX

I've been exploring the combination of Go and HTMX for building modern web applications, and I'm impressed by how much you can accomplish with minimal complexity.

## Why Go?

Go's simplicity and performance make it ideal for web servers. The standard library's net/http package provides everything you need to build robust web applications without additional frameworks.

## Why HTMX?

HTMX brings interactivity to your HTML without writing JavaScript. By adding attributes to your markup, you can create dynamic experiences that feel modern while keeping your server in control.

## The Benefits

- **Type safety**: Go's static typing catches errors at compile time
- **Performance**: Both Go and HTMX are incredibly fast
- **Simplicity**: Less JavaScript means less complexity
- **Progressive enhancement**: Works without JavaScript, better with it

This portfolio site is built entirely with this stack, and the development experience has been fantastic.`,
			publishedAt: time.Now().AddDate(0, 0, -7),
			tags:        []string{"Go", "HTMX", "Web Development"},
		},
		{
			title:   "The Power of Server-Side Rendering",
			slug:    "power-of-server-side-rendering",
			excerpt: "Why returning HTML from the server is making a comeback and how it simplifies modern web development.",
			content: `# The Power of Server-Side Rendering

After years of heavy client-side frameworks, there's a growing movement back to server-side rendering. Here's why it matters.

## Performance Benefits

Sending HTML from the server is fast. Really fast. Users see content immediately without waiting for JavaScript bundles to download and execute.

## Simplified Architecture

With SSR, your application logic lives in one place - the server. No need to duplicate validation, routing, and state management across client and server.

## Better SEO

Search engines can easily crawl and index your content when it's rendered on the server. No need for complex pre-rendering solutions.

## Accessibility

HTML-first development means your site works for everyone, including users with JavaScript disabled or using assistive technologies.

## HTMX Makes It Interactive

Tools like HTMX prove you don't need to sacrifice interactivity when choosing SSR. You get the best of both worlds.`,
			publishedAt: time.Now().AddDate(0, 0, -14),
			tags:        []string{"Web Development", "SSR", "Architecture"},
		},
		{
			title:   "Templ: Type-Safe HTML Templates in Go",
			slug:    "templ-type-safe-html-templates",
			excerpt: "How Templ brings compile-time safety and excellent developer experience to HTML templating in Go applications.",
			content: `# Templ: Type-Safe HTML Templates in Go

If you've worked with Go's html/template package, you know the pain of runtime template errors. Templ solves this by generating Go code from your templates.

## What is Templ?

Templ is a templating language that compiles to Go code. Write your HTML in .templ files, run the generator, and get type-safe Go functions.

## Key Benefits

### Compile-Time Checking

Syntax errors and type mismatches are caught when you run 'templ generate', not at runtime when a user visits your page.

### IDE Support

With the right extensions, you get autocomplete, syntax highlighting, and refactoring support in your templates.

### Component Composition

Build small, reusable components and compose them together. It feels like React, but generates Go code.

### Performance

Generated code is pure Go - no reflection, no runtime parsing. Just fast, efficient HTML generation.

## Example

You write your templates in .templ files with a syntax similar to JSX, and the generator creates Go functions you can call from your handlers. Type-safe, fast, and maintainable.

## Conclusion

Templ has transformed how I write Go web applications. If you're building with Go, give it a try.`,
			publishedAt: time.Now().AddDate(0, 0, -21),
			tags:        []string{"Go", "Templ", "Templates"},
		},
		{
			title:   "SQLite for Web Applications",
			slug:    "sqlite-for-web-applications",
			excerpt: "Why SQLite is an excellent choice for many web applications and how it simplifies deployment and development.",
			content: `# SQLite for Web Applications

SQLite often gets dismissed as a toy database, but it's actually a fantastic choice for many production web applications.

## Why SQLite?

### Single File Database

Your entire database is one file. Backup? Copy the file. Migration? Move the file. Simple.

### No Server Required

Unlike PostgreSQL or MySQL, SQLite runs in-process. No separate database server to manage, configure, or keep running.

### Performance

For read-heavy workloads (like most websites), SQLite is incredibly fast. It can handle thousands of requests per second.

### Reliability

SQLite is one of the most tested pieces of software in existence. It's battle-tested and rock-solid.

## Perfect Use Cases

- Portfolio sites (like this one!)
- Blogs and content sites
- Internal tools
- Mobile apps
- Prototypes and MVPs

## When to Choose PostgreSQL Instead

- High write concurrency
- Multiple application servers
- Need for advanced features (full-text search, JSON queries)
- Very large datasets (100GB+)

## Conclusion

Don't let the "lite" in SQLite fool you. For many applications, it's the perfect database choice.`,
			publishedAt: time.Now().AddDate(0, 0, -28),
			tags:        []string{"SQLite", "Databases", "Architecture"},
		},
		{
			title:   "Momentum-Based Animations with JavaScript",
			slug:    "momentum-based-animations-javascript",
			excerpt: "Creating smooth, physics-based animations using momentum and damping for natural-feeling user interfaces.",
			content: `# Momentum-Based Animations with JavaScript

The hero section of this portfolio uses momentum-based physics to create smooth, natural animations. Here's how it works.

## The Physics

Instead of directly setting positions, we calculate forces and apply momentum. Each frame, we update velocity based on forces, then update position based on velocity, and finally apply damping to create smooth deceleration.

## Why It Feels Better

Direct position updates feel robotic. Momentum creates smooth acceleration and deceleration that mimics real-world physics.

## Adding Interactivity

The hero grid responds to mouse position by applying forces to each square:

- **Attraction mode**: Squares gravitate toward cursor
- **Parting mode**: Hold mouse to make squares flee
- **Exponential scaling**: Forces increase exponentially for dramatic effect

## Easter Eggs

Rapid clicking accumulates "click energy" that triggers visual effects. Hidden interactions reward exploration.

## Performance

Use 'requestAnimationFrame' for smooth 60fps animations and only update elements that changed.

## Try It

Scroll up to the hero section and interact with the grid to feel the momentum physics in action!`,
			publishedAt: time.Now().AddDate(0, 0, -35),
			tags:        []string{"JavaScript", "Animation", "UX"},
		},
	}

	stmt, err := db.Prepare(`
		INSERT INTO blog_posts (title, slug, excerpt, content, published_at, tags)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, post := range posts {
		tagsJSON, err := json.Marshal(post.tags)
		if err != nil {
			return fmt.Errorf("marshal tags: %w", err)
		}

		_, err = stmt.Exec(
			post.title,
			post.slug,
			post.excerpt,
			post.content,
			post.publishedAt,
			string(tagsJSON),
		)
		if err != nil {
			return fmt.Errorf("insert post %s: %w", post.slug, err)
		}
	}

	log.Printf("Seeded %d blog posts", len(posts))
	return nil
}
