package handlers

import (
	"fmt"
	"net/http"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/export" {
		renderErrorPage(w, "Not Found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderErrorPage(w, fmt.Sprintf("Error parsing form data: %v", err), http.StatusBadRequest)
		return
	}

	output := r.FormValue("output")
	format := r.FormValue("format")

	if format != "txt" {
		renderErrorPage(w, "Unsupported format", http.StatusBadRequest)
		return
	}

	contentType := "text/plain"
	fileExtension := "txt"

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=ASCII-ART-WEB.%s", fileExtension))

	if output == "" {
		renderErrorPage(w, "No output provided", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(output)))
	_, err := w.Write([]byte(output))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
