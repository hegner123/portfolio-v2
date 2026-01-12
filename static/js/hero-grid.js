// Hero Grid Animation
document.addEventListener('DOMContentLoaded', function() {
    const heroGrid = document.getElementById('heroGrid');

    if (!heroGrid) {
        return;
    }

    // Grid configuration - responsive based on screen size
    function getGridConfig() {
        const width = window.innerWidth;

        // Mobile: <= 768px
        if (width <= 768) {
            return { ITEM_SIZE: 40, GAP: 6, PADDING: 6 };
        }
        // Tablet: 769px - 1024px
        else if (width <= 1024) {
            return { ITEM_SIZE: 50, GAP: 7, PADDING: 7 };
        }
        // Desktop: > 1024px
        else {
            return { ITEM_SIZE: 60, GAP: 8, PADDING: 8 };
        }
    }

    const itemMomentum = new Map();

    // Function to calculate and create grid items
    function createGridItems() {
        // Clear existing items and momentum
        heroGrid.innerHTML = '';
        itemMomentum.clear();

        // Get responsive grid configuration
        const config = getGridConfig();

        // Get grid dimensions
        const gridRect = heroGrid.getBoundingClientRect();
        const availableWidth = gridRect.width - (config.PADDING * 2);
        const availableHeight = gridRect.height - (config.PADDING * 2);

        // Calculate number of columns and rows needed to fill the space
        // Formula: floor((available + gap) / (size + gap))
        // This accounts for the fact that the last item doesn't need a gap after it
        const cols = Math.floor((availableWidth + config.GAP) / (config.ITEM_SIZE + config.GAP));
        const rows = Math.floor((availableHeight + config.GAP) / (config.ITEM_SIZE + config.GAP));
        const totalItems = cols * rows;

        // Create grid items with momentum tracking
        for (let i = 0; i < totalItems; i++) {
            const item = document.createElement('div');
            item.className = 'hero-grid-item';
            item.dataset.index = i;
            heroGrid.appendChild(item);

            // Initialize momentum for each item
            itemMomentum.set(item, { vx: 0, vy: 0, x: 0, y: 0 });
        }

        return document.querySelectorAll('.hero-grid-item');
    }

    // Initial grid creation
    let gridItems = createGridItems();

    // Animation state variables
    let mouseX = 0;
    let mouseY = 0;
    let isMouseDown = false;
    let partAmount = 0;
    let resetTimer = null;
    let squareless = false;

    // Track mouse/touch position globally on the document
    document.addEventListener('mousemove', function(e) {
        mouseX = e.clientX;
        mouseY = e.clientY;
    });

    // Touch support for mobile/tablet
    document.addEventListener('touchmove', function(e) {
        if (e.touches.length > 0) {
            mouseX = e.touches[0].clientX;
            mouseY = e.touches[0].clientY;
        }
    }, { passive: true });

    // Track mouse/touch down state globally on the hero section
    const heroSection = document.querySelector('.hero-section');
    if (heroSection) {
        // Mouse events
        heroSection.addEventListener('mousedown', function() {
            if (squareless) return;

            isMouseDown = true;
            partAmount++;

            // Clear existing reset timer
            if (resetTimer) {
                clearTimeout(resetTimer);
                resetTimer = null;
            }

            // Check for explosion state
            if (partAmount >= 10) {
                explodeSquares();
            }
        });

        heroSection.addEventListener('mouseup', function() {
            isMouseDown = false;

            // Start 0.5s timer to reset partAmount (only on hero section mouseup)
            if (!squareless && partAmount < 10) {
                resetTimer = setTimeout(function() {
                    partAmount = 0;
                    resetTimer = null;
                }, 500);
            }
        });

        // Touch events for mobile/tablet
        heroSection.addEventListener('touchstart', function(e) {
            if (squareless) return;

            isMouseDown = true;
            partAmount++;

            // Update position for touch
            if (e.touches.length > 0) {
                mouseX = e.touches[0].clientX;
                mouseY = e.touches[0].clientY;
            }

            // Clear existing reset timer
            if (resetTimer) {
                clearTimeout(resetTimer);
                resetTimer = null;
            }

            // Check for explosion state
            if (partAmount >= 10) {
                explodeSquares();
            }
        }, { passive: true });

        heroSection.addEventListener('touchend', function() {
            isMouseDown = false;

            // Start 0.5s timer to reset partAmount
            if (!squareless && partAmount < 10) {
                resetTimer = setTimeout(function() {
                    partAmount = 0;
                    resetTimer = null;
                }, 500);
            }
        }, { passive: true });
    }

    // Also track mouse/touch up outside the hero section (but don't start timer)
    document.addEventListener('mouseup', function() {
        isMouseDown = false;
    });

    document.addEventListener('touchend', function() {
        isMouseDown = false;
    }, { passive: true });

    // Explode squares off viewport
    function explodeSquares() {
        squareless = true;

        gridItems.forEach(item => {
            const rect = item.getBoundingClientRect();
            const itemCenterX = rect.left + rect.width / 2;
            const itemCenterY = rect.top + rect.height / 2;

            // Calculate direction from mouse
            const deltaX = itemCenterX - mouseX;
            const deltaY = itemCenterY - mouseY;
            const distance = Math.sqrt(deltaX * deltaX + deltaY * deltaY);

            // Normalize direction and apply massive force
            const normalizedX = deltaX / distance;
            const normalizedY = deltaY / distance;
            const explosionForce = 2000;

            item.style.transition = 'transform 2s ease-in-out, opacity 2s ease-out';
            item.style.transform = `translate(${normalizedX * explosionForce}px, ${normalizedY * explosionForce}px)`;
            item.style.opacity = '0';
        });

        // Delete squares after animation
        setTimeout(function() {
            heroGrid.innerHTML = '';
        }, 2000);
    }

    // Animation loop for gravitating effect with momentum
    function animateGrid() {
        if (squareless) {
            return;
        }

        gridItems.forEach(item => {
            const momentum = itemMomentum.get(item);
            if (!momentum) return;

            const rect = item.getBoundingClientRect();
            const itemCenterX = rect.left + rect.width / 2;
            const itemCenterY = rect.top + rect.height / 2;

            // Calculate distance from mouse
            const deltaX = mouseX - itemCenterX;
            const deltaY = mouseY - itemCenterY;
            const distance = Math.sqrt(deltaX * deltaX + deltaY * deltaY);

            // Maximum effect distance (constant, not scaled by partAmount)
            // Set to 1 square (responsive ITEM_SIZE)
            const config = getGridConfig();
            const maxDistance = config.ITEM_SIZE;
            const partMultiplier = Math.max(1, partAmount);

            let forceX = 0;
            let forceY = 0;

            if (distance < maxDistance) {
                // Calculate strength based on distance (closer = stronger)
                const strength = (maxDistance - distance) / maxDistance;

                // Base force amount (reduced for more fluid motion)
                let baseForce = 0.8;

                // Calculate force direction (gravitate towards mouse)
                forceX = (deltaX / distance) * strength * baseForce;
                forceY = (deltaY / distance) * strength * baseForce;

                // Apply partAmount multiplier to the force
                forceX *= partMultiplier;
                forceY *= partMultiplier;

                // If mouse is down, reverse the force (run away/part)
                if (isMouseDown) {
                    forceX = -forceX;
                    forceY = -forceY;
                }
            } else {
                // Add restoring force to pull back to origin when far from mouse
                const restoreStrength = 0.15;
                forceX = -momentum.x * restoreStrength;
                forceY = -momentum.y * restoreStrength;
            }

            // Apply force to velocity (momentum)
            momentum.vx += forceX;
            momentum.vy += forceY;

            // Apply damping/friction
            momentum.vx *= 0.85;
            momentum.vy *= 0.85;

            // Update position with velocity
            momentum.x += momentum.vx;
            momentum.y += momentum.vy;

            // Calculate maximum allowed offset from origin
            // baseAmount calculated to reach ~100px at 9 clicks (before explosion at 10)
            const baseAmount = 100 / 81; // ~1.234
            const maxOffset = baseAmount * (partMultiplier * partMultiplier);

            // Clamp position to not exceed max offset
            const currentDistance = Math.sqrt(momentum.x * momentum.x + momentum.y * momentum.y);
            if (currentDistance > maxOffset) {
                const scale = maxOffset / currentDistance;
                momentum.x *= scale;
                momentum.y *= scale;
                // Dampen velocity when hitting the limit
                momentum.vx *= 0.5;
                momentum.vy *= 0.5;
            }

            // Calculate opacity based on partAmount (0.03 base, up to 0.2 at 9 clicks)
            const baseOpacity = 0.03;
            const maxOpacity = 0.2;
            const opacityIncrease = (maxOpacity - baseOpacity) / 9; // Spread across 9 clicks
            const currentOpacity = Math.min(maxOpacity, baseOpacity + (opacityIncrease * partAmount));

            // Apply transform and opacity
            item.style.transform = `translate(${momentum.x}px, ${momentum.y}px)`;
            item.style.background = `rgba(255, 255, 255, ${currentOpacity})`;
        });

        requestAnimationFrame(animateGrid);
    }

    // Start animation
    animateGrid();

    // Recalculate grid on window resize
    let resizeTimeout;
    window.addEventListener('resize', function() {
        if (squareless) return;

        clearTimeout(resizeTimeout);
        resizeTimeout = setTimeout(function() {
            // Recreate grid with new dimensions
            gridItems = createGridItems();
        }, 250);
    });
});
