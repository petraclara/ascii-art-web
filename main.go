package main

import (
	"net/http"
	"html/template"
	"fmt"
	"strings"

	"ascii-art-web/asciiArt"
)

func main() {
	// Serve files inside templates/ (images, extra html if any)
	templateFS := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", templateFS))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

// homeHandler handles GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if r.URL.Path !="/"{
	http.Error(w,"Page not found",http.StatusNotFound)
	return
}
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}

// asciiArtHandler handles POST /ascii-art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Recover from panics → 500 error
	defer func() {
		if rec := recover(); rec != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()
if r.URL.Path !="/ascii-art"{
	http.Error(w,"Page not found",http.StatusNotFound)
	return
}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Missing input", http.StatusBadRequest)
		return
	}

	// Normalize newlines
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	lines := strings.Split(text, "\n")
	var finalResult string

	for _, line := range lines {
		if line == "" {
			finalResult += "\n"
			continue
		}

		part, err := asciiArt.AsciiArt(line, banner)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		finalResult += part + "\n"
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Prepare HTML-safe output
	fresult := strings.ReplaceAll(finalResult, " ", "&nbsp;")
	fresult = strings.ReplaceAll(fresult, "\n", "<br>")

	data := struct {
		Result template.HTML
	}{
		Result: template.HTML(fresult),
	}

	tmpl.Execute(w, data)
}
