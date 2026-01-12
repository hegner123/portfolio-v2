# Troubleshooting Guide

This document captures known issues, their root causes, and solutions discovered during development.

## Hero Grid Issues

### Grid Item Count Miscalculation on Mobile/Tablet

**Symptoms:**
- Grid squares not filling the viewport correctly on mobile/tablet
- Wrong number of grid items being created
- Grid layout appearing incorrect or misaligned

**Root Cause:**
Horizontal overflow on the page caused `getBoundingClientRect()` in the JavaScript to return incorrect dimensions for the hero grid container. When the page could scroll horizontally, the grid calculations based on `gridRect.width` and `gridRect.height` were using the wrong viewport dimensions.

**Solution:**
Add overflow protection to prevent horizontal scrolling:

```css
html {
    overflow-x: hidden;
    width: 100%;
}

body {
    overflow-x: hidden;
    width: 100%;
    max-width: 100vw;
}

.hero-section {
    width: 100%;
    max-width: 100vw;
    overflow: hidden;
}

section {
    width: 100%;
    max-width: 100vw;
    overflow-x: hidden;
}
```

**Key Insight:**
JavaScript calculations that depend on `getBoundingClientRect()` or viewport dimensions will be inaccurate if elements can overflow the viewport. Always ensure proper overflow control when doing dynamic layout calculations.

**Related Files:**
- `static/css/styles.css` - Overflow protection
- `static/js/hero-grid.js` - Grid calculation logic (lines 39-48)

---

## Navigation Issues

### Active Link Hover Text Disappearing

**Symptoms:**
- Active navigation link text becomes invisible when hovered
- Text appears to match the background color

**Root Cause:**
The general `.navigation__link:hover` rule applied a solid background (`rgba(255, 255, 255, 0.05)`), but active links use `-webkit-text-fill-color: transparent` to show a gradient through the text. The transparent text over the solid background made it invisible.

**Solution:**
Add specific hover styles for active links that maintain the gradient:

```css
.navigation__link--active:hover {
    background: var(--gradient-accent);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    transform: translateY(-2px);
}
```

**Related Files:**
- `static/css/navigation.css` (lines 102-108)

---

## Mobile Menu Issues

### Z-Index Stacking Problems

**Symptoms:**
- Mobile menu overlay appears behind page content
- Menu not properly covering elements when open

**Root Cause:**
Missing z-index on `.navigation__list` meant it didn't have a defined stacking context relative to other fixed/absolute positioned elements.

**Solution:**
Ensure proper z-index stacking order:
- Hamburger button: `z-index: 1001`
- Menu list: `z-index: 1000`
- Navigation bar: `z-index: 1000`
- Backdrop overlay: `z-index: 999`

**Related Files:**
- `static/css/navigation.css`

---

## Best Practices Learned

1. **Always control overflow** when doing JavaScript layout calculations based on viewport/container dimensions
2. **Test hover states** for elements with special text effects (gradients, transparency)
3. **Define z-index explicitly** for all fixed/absolute positioned elements that may overlap
4. **Use both `100vh` and `100dvh`** for mobile full-height elements to handle browser address bar changes
5. **Set explicit dimensions** (`width: 100%`, `max-width: 100vw`) to prevent unexpected overflow
