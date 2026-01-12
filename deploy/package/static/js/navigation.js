// Navigation functionality: active section detection, mobile menu, and scrolled state
document.addEventListener('DOMContentLoaded', function() {
    const navigation = document.querySelector('.navigation');
    const hamburger = document.querySelector('.navigation__hamburger');
    const navLinks = document.querySelectorAll('.navigation__link');
    const sections = document.querySelectorAll('section[id]');

    if (!navigation) return;

    // ========================================
    // Feature 0: Fix Navigation Links for Non-Home Pages
    // ========================================
    const isHomePage = window.location.pathname === '/';

    if (!isHomePage) {
        // On non-home pages, convert anchor links to full paths
        navLinks.forEach(link => {
            const href = link.getAttribute('href');
            if (href && href.startsWith('#')) {
                link.setAttribute('href', '/' + href);
            }
        });
    }

    // ========================================
    // Feature 1: Intersection Observer for Active Section Detection
    // ========================================
    if (sections.length > 0 && navLinks.length > 0) {
        const observerOptions = {
            threshold: 0.4,
            rootMargin: '-80px 0px -40% 0px'
        };

        const sectionObserver = new IntersectionObserver(function(entries) {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    const sectionId = entry.target.id;

                    // Remove active class from all links
                    navLinks.forEach(link => {
                        link.classList.remove('navigation__link--active');
                        link.removeAttribute('aria-current');
                    });

                    // Add active class to corresponding link
                    const activeLink = document.querySelector(`a[data-nav-link="${sectionId}"]`);
                    if (activeLink) {
                        activeLink.classList.add('navigation__link--active');
                        activeLink.setAttribute('aria-current', 'page');
                    }
                }
            });
        }, observerOptions);

        // Observe all sections
        sections.forEach(section => sectionObserver.observe(section));
    }

    // ========================================
    // Feature 2: Mobile Menu Toggle
    // ========================================
    if (hamburger) {
        hamburger.addEventListener('click', function() {
            const isExpanded = this.getAttribute('aria-expanded') === 'true';
            this.setAttribute('aria-expanded', !isExpanded);
            navigation.classList.toggle('navigation--menu-open');
        });

        // Close menu when any link is clicked
        navLinks.forEach(link => {
            link.addEventListener('click', function() {
                navigation.classList.remove('navigation--menu-open');
                hamburger.setAttribute('aria-expanded', 'false');
            });
        });

        // Close menu when clicking outside (on backdrop)
        navigation.addEventListener('click', function(e) {
            if (e.target === navigation && navigation.classList.contains('navigation--menu-open')) {
                navigation.classList.remove('navigation--menu-open');
                hamburger.setAttribute('aria-expanded', 'false');
            }
        });

        // Close menu on ESC key
        document.addEventListener('keydown', function(e) {
            if (e.key === 'Escape' && navigation.classList.contains('navigation--menu-open')) {
                navigation.classList.remove('navigation--menu-open');
                hamburger.setAttribute('aria-expanded', 'false');
                hamburger.focus();
            }
        });
    }

    // ========================================
    // Feature 3: Scrolled State Detection
    // ========================================
    function updateNavigationScrollState() {
        const scrollY = window.scrollY;

        if (scrollY > 100) {
            navigation.classList.add('navigation--scrolled');
        } else {
            navigation.classList.remove('navigation--scrolled');
        }
    }

    // Initial check
    updateNavigationScrollState();

    // Update on scroll (using requestAnimationFrame for performance)
    let scrollTimeout;
    window.addEventListener('scroll', function() {
        if (scrollTimeout) {
            window.cancelAnimationFrame(scrollTimeout);
        }
        scrollTimeout = window.requestAnimationFrame(updateNavigationScrollState);
    }, { passive: true });
});
