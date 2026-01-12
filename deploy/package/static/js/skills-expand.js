// Skills expand/collapse functionality
document.addEventListener('DOMContentLoaded', function() {
    const expandButtons = document.querySelectorAll('.skill-expand-btn');

    expandButtons.forEach(button => {
        button.addEventListener('click', function() {
            const category = this.closest('.skill-category');
            category.classList.toggle('expanded');
        });
    });
});
