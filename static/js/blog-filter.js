// Blog filter active state management
document.addEventListener('DOMContentLoaded', () => {
    const filterContainer = document.querySelector('.blog-feed__filters');

    if (!filterContainer) return;

    // Handle filter button clicks
    filterContainer.addEventListener('click', (event) => {
        const button = event.target.closest('.blog-feed__filter-tag');

        if (!button) return;

        // Remove active class from all buttons
        const allButtons = filterContainer.querySelectorAll('.blog-feed__filter-tag');
        allButtons.forEach(btn => btn.classList.remove('blog-feed__filter-tag--active'));

        // Add active class to clicked button
        button.classList.add('blog-feed__filter-tag--active');
    });
});
