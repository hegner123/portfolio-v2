# Design System Guide

This guide documents the design system for Portfolio V2, defined in `/static/css/design-system.css`.

## Overview

The design system provides a centralized set of design tokens (CSS custom properties) that ensure consistency across all components. All component CSS files should reference these tokens rather than using hard-coded values.

## Color Palette

### Background Colors
- `--color-bg-primary`: #000000 - Primary background (pure black)
- `--color-bg-secondary`: #0a0a0a - Secondary background (near black)
- `--color-surface-elevated`: #1a1a1a - Elevated surfaces (cards, modals)
- `--color-surface-hover`: #222222 - Hover state for surfaces

### Text Colors
- `--color-text-primary`: #ffffff - Primary text (pure white)
- `--color-text-secondary`: #e0e0e0 - Secondary text (light gray)
- `--color-text-tertiary`: #d0d0d0 - Tertiary text (medium gray)
- `--color-text-muted`: #a0a0a0 - Muted text (dark gray)

### Accent Colors
- `--color-accent-purple`: #764ba2 - Primary purple accent
- `--color-accent-blue`: #667eea - Primary blue accent
- `--color-accent-purple-light`: #9575cd - Light purple variant
- `--color-accent-blue-light`: #8b9ff3 - Light blue variant

### Status Colors
- `--color-success`: #10b981 - Success state (green)
- `--color-warning`: #f59e0b - Warning state (orange)
- `--color-error`: #ef4444 - Error state (red)
- `--color-info`: #3b82f6 - Info state (blue)

### Gradients
- `--gradient-primary`: Linear gradient (blue to purple, 135deg)
- `--gradient-accent`: Linear gradient (blue to purple, 90deg)
- `--gradient-accent-reverse`: Linear gradient (purple to blue, 90deg)
- `--gradient-bg-dark`: Dark background gradient
- `--gradient-radial-glow`: Radial glow effect (blue)
- `--gradient-shimmer`: Shimmer animation gradient

## Typography

### Font Families
- `--font-family-base`: System font stack (San Francisco, Segoe UI, etc.)
- `--font-family-heading`: Same as base (can be customized)
- `--font-family-mono`: Monospace font stack (SF Mono, Monaco, etc.)

### Font Sizes (Type Scale)
| Token | Size | Pixels | Usage |
|-------|------|--------|-------|
| `--font-size-xs` | 0.75rem | 12px | Small labels, captions |
| `--font-size-sm` | 0.875rem | 14px | Small text, metadata |
| `--font-size-base` | 1rem | 16px | Body text (base) |
| `--font-size-md` | 1.125rem | 18px | Large body text |
| `--font-size-lg` | 1.25rem | 20px | Small headings |
| `--font-size-xl` | 1.5rem | 24px | H4 headings |
| `--font-size-2xl` | 1.75rem | 28px | H3 headings |
| `--font-size-3xl` | 2rem | 32px | H2 headings |
| `--font-size-4xl` | 2.5rem | 40px | H1 headings |
| `--font-size-5xl` | 3rem | 48px | Large hero text |
| `--font-size-6xl` | 3.5rem | 56px | Extra large hero text |
| `--font-size-7xl` | 4rem | 64px | Display text |

### Font Weights
- `--font-weight-light`: 300
- `--font-weight-normal`: 400
- `--font-weight-medium`: 500
- `--font-weight-semibold`: 600
- `--font-weight-bold`: 700
- `--font-weight-extrabold`: 800

### Line Heights
- `--line-height-tight`: 1.25 - Headings
- `--line-height-snug`: 1.375 - Tight text blocks
- `--line-height-normal`: 1.5 - Standard body text
- `--line-height-relaxed`: 1.625 - Comfortable reading
- `--line-height-loose`: 1.8 - Very spacious text

### Letter Spacing
- `--letter-spacing-tight`: -0.025em
- `--letter-spacing-normal`: 0
- `--letter-spacing-wide`: 0.025em
- `--letter-spacing-wider`: 0.05em
- `--letter-spacing-widest`: 0.1em - Uppercase labels

## Spacing

### Spacing Scale
| Token | Size | Pixels | Usage |
|-------|------|--------|-------|
| `--spacing-xs` | 0.25rem | 4px | Minimal spacing |
| `--spacing-sm` | 0.5rem | 8px | Small gaps |
| `--spacing-md` | 1rem | 16px | Standard spacing |
| `--spacing-lg` | 1.5rem | 24px | Medium spacing |
| `--spacing-xl` | 2rem | 32px | Large spacing |
| `--spacing-2xl` | 2.5rem | 40px | Extra large spacing |
| `--spacing-3xl` | 3rem | 48px | Section padding |
| `--spacing-4xl` | 4rem | 64px | Large section padding |
| `--spacing-5xl` | 6rem | 96px | Hero section padding |
| `--spacing-6xl` | 8rem | 128px | Maximum spacing |

### Section Padding
- `--section-padding-mobile`: 3rem (48px)
- `--section-padding-tablet`: 4rem (64px)
- `--section-padding-desktop`: 6rem (96px)

## Layout & Sizing

### Container Widths
- `--container-sm`: 640px
- `--container-md`: 768px
- `--container-lg`: 1024px
- `--container-xl`: 1200px - Default max-width
- `--container-2xl`: 1400px

### Responsive Breakpoints
**Note:** CSS custom properties cannot be used directly in media queries. Use these values:

- **Mobile (xs)**: 480px and below
- **Small (sm)**: 640px
- **Tablet (md)**: 768px
- **Desktop (lg)**: 1024px
- **Large (xl)**: 1280px

### Border Radius
- `--border-radius-sm`: 6px - Buttons, tags
- `--border-radius-md`: 8px - Inputs
- `--border-radius-lg`: 12px - Cards (small)
- `--border-radius-xl`: 16px - Cards (large)
- `--border-radius-2xl`: 20px - Prominent cards
- `--border-radius-full`: 9999px - Pills, circles

### Z-Index Scale
Use these standardized z-index values to prevent conflicts:

- `--z-index-dropdown`: 100
- `--z-index-sticky`: 200 - Sticky navigation
- `--z-index-fixed`: 300
- `--z-index-modal-backdrop`: 400
- `--z-index-modal`: 500
- `--z-index-popover`: 600
- `--z-index-tooltip`: 700

## Borders & Shadows

### Borders
- `--border-subtle`: 1px solid rgba(255, 255, 255, 0.05)
- `--border-soft`: 1px solid rgba(255, 255, 255, 0.1)
- `--border-accent`: 1px solid rgba(102, 126, 234, 0.3)
- `--border-accent-strong`: 1px solid rgba(102, 126, 234, 0.5)

### Shadows
- `--shadow-xs`: Minimal elevation
- `--shadow-sm`: Low elevation
- `--shadow-md`: Medium elevation
- `--shadow-lg`: High elevation
- `--shadow-xl`: Very high elevation
- `--shadow-2xl`: Maximum elevation

### Glow Effects
- `--shadow-glow-sm`: Small glow (20px blur)
- `--shadow-glow-md`: Medium glow (40px blur)
- `--shadow-glow-lg`: Large glow (60px blur)
- `--shadow-accent`: Combined blue/purple glow
- `--shadow-accent-strong`: Intense combined glow

## Transitions & Animations

### Timing Functions
- `--ease-linear`: linear
- `--ease-in`: cubic-bezier(0.4, 0, 1, 1)
- `--ease-out`: cubic-bezier(0, 0, 0.2, 1)
- `--ease-in-out`: cubic-bezier(0.4, 0, 0.2, 1)
- `--ease-bounce`: cubic-bezier(0.68, -0.55, 0.265, 1.55)
- `--ease-smooth`: cubic-bezier(0.25, 0.46, 0.45, 0.94)

### Durations
- `--duration-instant`: 100ms - Instant feedback
- `--duration-fast`: 200ms - Quick transitions
- `--duration-normal`: 300ms - Standard transitions
- `--duration-slow`: 400ms - Deliberate transitions
- `--duration-slower`: 500ms - Slow, noticeable transitions

### Common Transitions
- `--transition-base`: All properties, normal speed
- `--transition-fast`: All properties, fast speed
- `--transition-slow`: All properties, slow speed
- `--transition-colors`: Background, text, border colors
- `--transition-transform`: Transform property only
- `--transition-opacity`: Opacity property only

### Keyframe Animations

Available animations:
- `fadeIn` / `fadeOut`
- `slideInUp` / `slideInDown` / `slideInLeft` / `slideInRight`
- `scaleIn`
- `pulse`
- `spin`
- `shimmer`
- `gradientShift`

Usage:
```css
.element {
    animation: fadeIn var(--duration-normal) var(--ease-in-out);
}
```

## Effects & Utilities

### Backdrop Blur
- `--blur-sm`: 4px
- `--blur-md`: 10px
- `--blur-lg`: 20px
- `--blur-xl`: 40px

Usage:
```css
backdrop-filter: blur(var(--blur-md));
```

### Focus States

All interactive elements automatically receive focus styling via the design system:

```css
:focus-visible {
    outline: 2px solid var(--color-accent-blue);
    outline-offset: 2px;
}
```

Override if needed:
```css
.custom-button:focus-visible {
    outline-color: var(--color-accent-purple);
    outline-offset: 4px;
}
```

## Accessibility Features

### Reduced Motion
The design system respects user preferences for reduced motion:

```css
@media (prefers-reduced-motion: reduce) {
    * {
        animation-duration: 0.01ms !important;
        transition-duration: 0.01ms !important;
        scroll-behavior: auto !important;
    }
}
```

### Focus Management
- `:focus-visible` is used instead of `:focus` for keyboard-only focus indicators
- Mouse clicks don't show outlines (better UX)
- Keyboard navigation shows clear focus indicators (accessibility)

### Print Styles
Print-friendly styles are automatically applied when printing:
- Dark backgrounds removed
- Colors converted to print-safe values
- Shadows and effects removed

## Usage Examples

### Using Color Tokens
```css
.component {
    background: var(--color-surface-elevated);
    color: var(--color-text-secondary);
    border: var(--border-accent);
}
```

### Using Spacing
```css
.section {
    padding: var(--section-padding-desktop);
    margin-bottom: var(--spacing-4xl);
}

@media (max-width: 768px) {
    .section {
        padding: var(--section-padding-mobile);
        margin-bottom: var(--spacing-2xl);
    }
}
```

### Using Typography
```css
.heading {
    font-size: var(--font-size-3xl);
    font-weight: var(--font-weight-bold);
    line-height: var(--line-height-tight);
    letter-spacing: var(--letter-spacing-tight);
}
```

### Using Transitions
```css
.button {
    transition: var(--transition-base);
    border-radius: var(--border-radius-md);
}

.button:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-accent);
}
```

### Using Gradients
```css
.hero-heading {
    background: var(--gradient-primary);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}
```

## Best Practices

### DO:
- ✅ Use design tokens for all values
- ✅ Reference existing gradients and shadows
- ✅ Use standardized spacing scale
- ✅ Follow the established type scale
- ✅ Use semantic naming (e.g., `--color-text-secondary` not `--color-gray-300`)
- ✅ Add vendor prefixes for experimental CSS features
- ✅ Test with reduced motion preferences
- ✅ Ensure proper focus states on all interactive elements

### DON'T:
- ❌ Hard-code color values (e.g., `#667eea`)
- ❌ Use arbitrary spacing values (e.g., `13px`)
- ❌ Create one-off gradients or shadows
- ❌ Skip vendor prefixes for `background-clip`, `backdrop-filter`, etc.
- ❌ Remove focus indicators without providing alternatives
- ❌ Use CSS custom properties in media query conditions (they don't work there)

## Browser Support

The design system is designed to work in:
- Chrome/Edge (latest 2 versions)
- Firefox (latest 2 versions)
- Safari (latest 2 versions)
- iOS Safari (latest 2 versions)

Fallbacks are provided for:
- `backdrop-filter` (falls back to solid background)
- Gradient text (falls back to solid color)
- CSS custom properties (not supported in IE11, which is acceptable for modern web apps)

## Maintenance

When updating the design system:

1. Update `/static/css/design-system.css`
2. Update this documentation
3. Test changes across all components
4. Ensure no hard-coded values remain in component files
5. Verify accessibility and responsive behavior
6. Update IMPLEMENTATION.md if adding new features

## Migration from Hard-Coded Values

If you find hard-coded values in component CSS:

```css
/* Before */
.card {
    color: #e0e0e0;
    padding: 32px;
    border-radius: 12px;
}

/* After */
.card {
    color: var(--color-text-secondary);
    padding: var(--spacing-xl);
    border-radius: var(--border-radius-lg);
}
```

This ensures:
- Consistency across the site
- Easy theme updates (change once, apply everywhere)
- Better maintainability
- Self-documenting code (semantic names)
