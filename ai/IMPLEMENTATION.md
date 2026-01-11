# IMPLEMENTATION PLAN

## Project Overview
Single-page portfolio website with:
- Hero section with eye-catching animation
- About section
- Blog feed
- Project feed
- Contact section

**Tech Stack**: Go + Templ + HTMX

---

## Status: In Progress

**Current Phase**: Design Polish & Accessibility (Phase 9)
**Next Action**: Review and enhance design consistency, accessibility, and cross-browser compatibility

**Last Updated**: Phases 7-8 completed - 2026-01-11

---

## Implementation Phases

### Phase 1: Project Foundation
**Status**: ✅ Completed

- [x] Initialize Go module
- [x] Install Templ dependency
- [x] Install Templ CLI tool
- [x] Create directory structure (handlers/, templates/, static/)
- [x] Download HTMX (or link to CDN)
- [x] Create basic main.go with server setup

---

### Phase 2: Core Application Structure
**Status**: ✅ Completed

- [x] Create base template/layout component
- [x] Set up static file server for CSS/JS/images
- [x] Create main route handler
- [x] Test basic server functionality

---

### Phase 3: Hero Section
**Status**: ✅ Completed

**Requirements**:
- Eye-catching animation (interactive grid with momentum physics)
- Name/title
- Brief tagline or introduction

**Tasks**:
- [x] Design hero animation concept
- [x] Create Hero templ component
- [x] Implement animation (CSS and JS with momentum physics)
- [x] Style hero section
- [x] Make responsive

**Implementation Details**:
- Interactive grid of semi-transparent squares
- Squares gravitate towards mouse (attraction mode)
- Hold mouse down to make squares run away (parting mode)
- Click accumulation system with 0.5s timer
- Momentum-based physics with damping
- Exponential force scaling (squared partAmount)
- Easter egg: 10 rapid clicks triggers explosion
- Dynamic opacity based on click count
- Accessible (aria-hidden on decorative elements)

---

### Phase 4: About Section
**Status**: ✅ Completed

**Requirements**:
- Personal/professional bio
- Skills or technologies
- Photo (optional)

**Tasks**:
- [x] Create About templ component
- [x] Add content structure (bio and skills grid)
- [x] Style about section
- [x] Make responsive

**Implementation Details**:
- Bio section with personalized content about pragmatic problem-solving approach
- Skills grid organized into 9 comprehensive categories:
  - Backend (9 technologies)
  - Frontend (8 technologies)
  - Databases (6 platforms)
  - AI & LLM Integration (6 skills)
  - Cloud & Infrastructure (5 platforms)
  - CMS Platforms (4 platforms)
  - Tools & Platforms (4 tools)
  - Security & Compliance (4 items)
  - Practices (8 methodologies)
- Expand/collapse functionality showing top 3 skills per category
- Skill count indicators next to category headings
- Dark theme with CSS variables for maintainability
- Skills card with prominent purple/blue glow effect
- Gradient background and card-based design
- Decorative gradient underline on heading
- Fully responsive with single-column layout on mobile
- Interactive JavaScript for smooth expand/collapse
- Accessible semantic HTML structure with ARIA labels

---

### Phase 5: Blog Feed
**Status**: ✅ Completed

**Requirements**:
- Display list of blog posts
- Post preview (title, date, excerpt)
- Link to full post or expand inline
- Dynamic loading with HTMX

**Tasks**:
- [x] Design blog post data structure
- [x] Create blog storage/retrieval (SQLite database)
- [x] Create BlogFeed templ component
- [x] Create BlogPost templ component
- [x] Create handler for blog list endpoint
- [x] Create handler for individual blog post
- [x] Implement HTMX loading (Load More button)
- [x] Style blog feed
- [x] Make responsive

**Implementation Details**:
- SQLite database with blog_posts table
- BlogPost and BlogPostPreview models
- Database layer with GetBlogPosts, GetBlogPostBySlug, CountBlogPosts functions
- 5 seed blog posts about Go, HTMX, Templ, SQLite, and JavaScript animations
- BlogFeed component with heading and post container
- BlogPostCard component showing date, title, excerpt, and tags
- BlogPostList component for HTMX responses
- Load More button showing 3 posts initially
- API endpoint at /api/blog/posts for pagination
- Dark theme styling with purple/blue gradients matching About section
- Card hover effects with glow and transform
- Gradient text for post titles
- Tag badges with purple theme
- Fully responsive with mobile breakpoints at 768px and 480px
- HTMX loading states for better UX

---

### Phase 6: Project Feed
**Status**: ✅ Completed

**Requirements**:
- Display list of projects
- Project preview (title, description, technologies, image)
- Link to project details or external site
- Dynamic loading with HTMX

**Tasks**:
- [x] Design project data structure
- [x] Create project storage/retrieval (SQLite database)
- [x] Create ProjectFeed templ component
- [x] Create ProjectCard templ component
- [x] Create handler for project list endpoint
- [x] Implement HTMX loading (grid view with Load More)
- [x] Style project feed
- [x] Make responsive

**Implementation Details**:
- SQLite database with projects table
- Project model with Title, Description, Technologies, GithubURL, ImageURL, Featured flag
- Database layer with GetProjects, GetProjectBySlug, CountProjects functions
- 6 seed projects (3 featured, 3 regular) with diverse tech stacks
- ProjectFeed component with heading and grid container
- ProjectCard component with image, featured badge, description, tech tags, and GitHub link
- ProjectList component for HTMX responses
- Load More button showing 3 projects initially
- API endpoint at /api/projects for pagination
- Dark theme styling with purple/blue gradients matching blog/about sections
- Grid layout (auto-fill minmax 350px) with responsive breakpoints
- Card hover effects with glow, transform, and image zoom
- Featured badge overlay on project images
- GitHub icon SVG with external link attributes
- Fully responsive with mobile breakpoints at 768px and 480px
- HTMX loading states for better UX

---

### Phase 7: Contact Section
**Status**: ✅ Completed

**Requirements**:
- Contact information display
- Links to email and social profiles
- Responsive design matching site theme

**Tasks**:
- [x] Create ContactForm templ component
- [x] Add email contact (mailto link)
- [x] Add LinkedIn profile link
- [x] Add GitHub profile link
- [x] Style contact section
- [x] Make responsive

**Implementation Details**:
- **Contact Info Display**: Clean layout with email, LinkedIn, and GitHub links
- **SVG Icons**: Email envelope, LinkedIn logo, GitHub logo (inline SVG)
- **Styling**: Dark theme with purple/blue gradient accents matching site design
- **Links**:
  - Email: mailto:hegner123@gmail.com
  - LinkedIn: linkedin.com/in/michaelhegner (opens in new tab)
  - GitHub: github.com/hegner123 (opens in new tab)
- **Accessibility**: External links have rel="noopener noreferrer"
- **Responsive**: Mobile-optimized layout
- **Design**: Gradient heading underline, icon-based layout, hover effects on links

---

### Phase 8: Navigation & Smooth Scrolling
**Status**: ✅ Completed

**Requirements**:
- Sticky navigation or scroll-to-section links
- Smooth scrolling between sections
- Active section highlighting

**Tasks**:
- [x] Create Navigation templ component
- [x] Implement smooth scroll (CSS or JS)
- [x] Add section anchors
- [x] Highlight active section on scroll
- [x] Style navigation
- [x] Make responsive

**Implementation Details**:
- **Sticky Navigation**: Glassmorphic navigation bar with backdrop blur
- **Smooth Scrolling**: CSS `scroll-behavior: smooth` with `scroll-padding-top` for nav offset
- **Active Section Detection**: Intersection Observer API with 40% threshold
- **Mobile Navigation**: Hamburger menu with slide-in panel (≤768px)
- **Styling**: Purple/blue gradient accents, hover effects (translateY -2px), scrolled state with glow
- **Responsive**: Desktop horizontal layout, mobile hamburger menu
- **Accessibility**: Keyboard navigation, ARIA labels, ESC to close, focus indicators
- **JavaScript**: Three features - active section highlighting, mobile menu toggle, scrolled state detection
- **Performance**: RequestAnimationFrame for scroll updates, passive event listeners
- **Files Created**: navigation.templ, navigation.css, navigation.js
- **Files Modified**: layout.templ (integration), hero.templ (id="home"), styles.css (smooth scroll)

---

### Phase 9: Styling & Design Polish
**Status**: Not Started

**Tasks**:
- [ ] Create main CSS file with design system
- [ ] Define color scheme
- [ ] Define typography
- [ ] Add animations and transitions
- [ ] Ensure mobile responsiveness across all sections
- [ ] Cross-browser testing
- [ ] Accessibility improvements (ARIA labels, keyboard nav)

---

### Phase 10: Performance & Optimization
**Status**: Not Started

**Tasks**:
- [ ] Optimize images
- [ ] Minify CSS
- [ ] Implement lazy loading for images
- [ ] Test page load performance
- [ ] Add meta tags for SEO
- [ ] Add Open Graph tags for social sharing

---

### Phase 11: Deployment
**Status**: Not Started

**Tasks**:
- [ ] Choose hosting platform
- [ ] Set up production environment
- [ ] Configure domain (if applicable)
- [ ] Deploy application
- [ ] Test production deployment
- [ ] Set up monitoring/logging

---

## Design Decisions to Make

1. ~~**Hero Animation**: What type of animation?~~ ✅ **Decided**: Interactive grid with momentum physics
2. ~~**Blog Storage**: File-based markdown? Database? Headless CMS?~~ ✅ **Decided**: SQLite database with seeded posts
3. ~~**Project Storage**: File-based? Database? GitHub API integration?~~ ✅ **Decided**: SQLite database with seeded projects
4. **Contact Form**: Email service (SMTP, SendGrid, etc.)? Store in database?
5. **Color Scheme**: Dark mode? Light mode? Theme switcher?
6. **Fonts**: Custom fonts or system fonts?

---

## Current Blockers
None

---

## Notes
- Single-page application means all sections load on initial page load
- HTMX will be used for dynamic content loading within feeds
- Templ generates type-safe Go code from templates
- Remember to run `templ generate` after modifying .templ files
