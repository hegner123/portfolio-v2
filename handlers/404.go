package handlers

import (
	"net/http"

	"portfolio-v2/templates"
)

// NotFoundHandler handles 404 errors with a custom page
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	component := templates.NotFound()
	component.Render(r.Context(), w)
}
