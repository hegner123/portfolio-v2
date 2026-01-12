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

## Status: Complete

**Current Phase**: Production & Maintenance
**Next Action**: Monitor application, add features as needed

**Last Updated**: Phase 11 completed - 2026-01-12

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
**Status**: ✅ Completed

**Tasks**:
- [x] Create main CSS file with design system
- [x] Define color scheme
- [x] Define typography
- [x] Add animations and transitions
- [x] Ensure mobile responsiveness across all sections
- [x] Cross-browser testing (vendor prefixes added)
- [x] Accessibility improvements (ARIA labels, keyboard nav)

**Implementation Details**:
- **Design System Created**: `/static/css/design-system.css` - comprehensive design tokens
  - Color palette: Background, text, accent, status colors with gradients
  - Typography: Type scale, font families, weights, line heights, letter spacing
  - Spacing: Standardized scale from 4px to 128px
  - Layout: Container widths, breakpoints, border radius, z-index scale
  - Borders & Shadows: Consistent borders and glow effects
  - Transitions: Timing functions, durations, pre-built transitions
  - Animations: 10+ keyframe animations (fadeIn, slideIn, pulse, shimmer, etc.)
  - Accessibility: Focus states, reduced motion support, print styles
- **Responsive Breakpoints Standardized**: All components now use consistent breakpoints
  - Mobile (xs): 480px and below
  - Tablet (md): 768px and below  - Desktop (lg): 1024px and above
- **Accessibility Enhanced**:
  - Added `aria-labelledby` to all section elements
  - Added `aria-hidden="true"` to decorative elements
  - Added `aria-label` to all interactive buttons
  - Added `role="group"` and `aria-pressed` to filter buttons
  - Global `:focus-visible` styles for keyboard navigation
  - Proper semantic HTML throughout (section, article, time, nav)
- **Vendor Prefixes Added**:
  - `-webkit-user-select` and `-moz-user-select` for user interaction control
  - `-webkit-background-clip` and `-webkit-text-fill-color` for gradient text (already present)
  - `backdrop-filter` fallbacks in navigation component
- **Transitions Optimized**:
  - All components updated to use design system transition variables
  - Consistent animation timing across the site
  - Performance-optimized with `will-change` where appropriate
- **Documentation Created**: `/ai/design-system-guide.md` - comprehensive design system guide
  - Complete token reference
  - Usage examples
  - Best practices
  - Migration guide
  - Browser support information

---

### Phase 10: Performance & Optimization
**Status**: ✅ Completed

**Tasks**:
- [x] Optimize images
- [x] Minify CSS (build script created)
- [x] Implement lazy loading for images
- [x] Add performance optimization headers
- [x] Add meta tags for SEO
- [x] Add Open Graph tags for social sharing

**Implementation Details**:
- **SEO Meta Tags Added** to `layout.templ`:
  - Description, keywords, author, robots directives
  - Language and revisit-after tags for crawler optimization
  - Comprehensive metadata for search engine indexing
- **Open Graph Tags** for social media sharing:
  - og:type, og:title, og:description, og:site_name, og:locale
  - Optimized for Facebook, LinkedIn, and other OG-compatible platforms
- **Twitter Card Tags**:
  - summary_large_image card type
  - Title and description optimized for Twitter sharing
- **Performance Optimizations**:
  - **Preconnect/DNS-prefetch**: Added for unpkg.com (HTMX CDN)
  - **Resource Preload**: Critical CSS files preloaded (design-system.css, styles.css)
  - **Script Optimization**: HTMX loaded with `defer`, `integrity`, and `crossorigin` attributes
  - **Lazy Loading**: Verified on all project card images (`loading="lazy"`)
  - **Theme Colors**: Set for mobile browsers (theme-color, msapplication-TileColor)
- **Build Scripts Created** (`/scripts/`):
  - **build-prod.sh**: Production build with optimizations
    - Generates Templ templates
    - Tidies Go modules
    - Builds binary with `-ldflags="-s -w"` (20-30% size reduction)
    - Auto-detects and minifies CSS (csso-cli or clean-css-cli)
    - Optional deployment tarball creation
    - Color-coded output with progress indicators
  - **build-dev.sh**: Quick development build
    - Generates templates
    - Builds binary without optimizations
    - Fast iteration for development
  - **README.md**: Comprehensive documentation
    - Usage instructions for both scripts
    - CSS minification setup guide
    - Deployment workflows (manual and tarball)
    - Troubleshooting section
    - CI/CD integration examples
- **Image Optimization**:
  - Lazy loading already implemented on project cards
  - No static images currently (placeholder URLs used)
  - Build script ready for future image optimization
- **Performance Monitoring Ready**:
  - All meta tags in place for analytics integration
  - Proper semantic HTML for better indexing
  - Accessibility features enhance SEO ranking

---

### Phase 11: Deployment
**Status**: ✅ Completed

**Tasks**:
- [x] Choose hosting platform (Linode)
- [x] Set up production environment
- [x] Configure domain (michaelhegner.com)
- [x] Deploy application
- [x] Test production deployment
- [x] Set up monitoring/logging

**Implementation Details**:
- **Server**: Linode VPS (50.116.26.167)
- **Domain**: michaelhegner.com with automatic HTTPS
- **Web Server**: Caddy reverse proxy with Let's Encrypt
- **Application**: Go binary as systemd service on port 8080
- **Database**: SQLite with pure Go driver (modernc.org/sqlite) for CGO-free cross-compilation
- **Deployment Scripts**: Complete automation in `/deploy` directory
  - `deploy.sh` - Application deployment with rsync
  - `update-caddy.sh` - Caddy configuration deployment
  - `full-deploy.sh` - Complete deployment pipeline
- **Deployment Features**:
  - Cross-compilation: Mac → Linux (CGO_ENABLED=0 GOOS=linux GOARCH=amd64)
  - Fast transfers: rsync with --delete for incremental updates
  - SSH key authentication: passwordless deployments
  - Environment-based configuration: .env file for admin credentials
  - Passwordless sudo: systemctl commands automated
  - Service management: systemd for auto-restart and lifecycle
- **Security**:
  - HTTPS with automatic certificate renewal
  - Security headers (HSTS, X-Frame-Options, CSP)
  - HTTP Basic Auth for admin routes
  - Passwordless sudo limited to specific commands
- **Monitoring**: System logs via journalctl for both Caddy and portfolio services
- **Bug Fixes Deployed**:
  - Blog edit form URL generation (fmt.Sprintf for proper ID conversion)
  - Featured badge text color (-webkit-text-fill-color for Safari)

**Deployment Workflow**:
1. Update code locally
2. Run `./deploy/deploy.sh` (reads password from .env)
3. Builds Linux binary, syncs via rsync, restarts service
4. Live at https://michaelhegner.com in ~30 seconds

---

### Phase 12: Admin Security
**Status**: Phase 1 Complete ✅

**Tasks**:
- [x] Choose authentication method (Phase 1 - HTTP Basic Auth)
- [x] Implement authentication middleware
- [x] Secure `/admin/blog/new` route
- [x] Secure `/admin/project/new` route
- [ ] Add login/logout functionality (Phase 2 - Optional)
- [ ] Add CSRF protection (Phase 2 - Optional)
- [ ] Implement rate limiting (Phase 2 - Optional)
- [ ] Add audit logging (Phase 2 - Optional)
- [x] Test security measures

**Implementation Details**:
- **Method**: HTTP Basic Authentication (Phase 1)
- **Protected Routes**:
  - `/admin/blog/new` - Create blog posts
  - `/admin/project/new` - Create projects
- **Credentials**: Set via environment variables or .env file
  - `ADMIN_USERNAME=admin`
  - `ADMIN_PASSWORD` - From .env file or prompted during deployment
- **Files Created**:
  - `middleware/auth.go` - Basic auth middleware
  - `deploy/ADMIN-AUTH-GUIDE.md` - User guide
  - `.env.example` - Environment configuration template
- **Security Features**:
  - Password validation (10+ characters)
  - Failed login attempt logging
  - HTTPS encryption (via Caddy)
  - No password stored in git (.env in .gitignore)

**Usage**:
- Visit admin pages → Browser prompts for credentials
- Username: `admin`
- Password: Set during deployment

**Upgrade Path**:
See `/ai/ADMIN-SECURITY-IMPLEMENTATION.md` for Phase 2 (Session-Based Auth) upgrade guide

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
