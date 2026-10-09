package main

import (
	"net/http"
)

func main() {
	// Serve your HTML, JavaScript, and asset folders
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Explicit handler to serve the webmanifest with the required Content-Type
	http.HandleFunc("/manifest.webmanifest", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/manifest+json")
		http.ServeFile(w, r, "./static/manifest.webmanifest")
	})

	http.ListenAndServe(":8080", nil)
}
