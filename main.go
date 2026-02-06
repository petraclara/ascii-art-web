package main

import (
	"net/http"
	"html/template"
	"fmt"
	"strings"
)

func main(){
		// Serve files inside templates/ (images, extra html if any)
	templateFS := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", templateFS))

	http.HandleFunc("/",  homeHandler)
	http.HandleFunc("/ascii-art",  asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
//homeHandler
func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}
		tmpl,err := template.ParseFiles("templates/index.html")
		if err != nil{
			http.Error(w, "Template not found", http.StatusNotFound)
			return

		}
		tmpl.Execute(w, nil)

	}

	//asciiArtHandler
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/ascii-art"{
		http.NotFound(w,r)
		return
	}
	
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Missing input", http.StatusBadRequest)
		return
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	// Split input into lines and generate ASCII art line by line
	lines := strings.Split(text, "\n")
	var finalResult string
	for _, line := range lines {
		if line == "" {
			finalResult += "\n"
			continue
		}

		part, err := AsciiArt(line, banner)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		finalResult += part + "\n"
	}

	// Parse the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Convert spaces and newlines for HTML display
	fresult := strings.ReplaceAll(finalResult, " ", "&nbsp;")
	fresult = strings.ReplaceAll(fresult, "\n", "<br>")

	// Pass to template
	data := struct {
		Result template.HTML
	}{
		Result: template.HTML(fresult),
	}

	tmpl.Execute(w, data)
}
