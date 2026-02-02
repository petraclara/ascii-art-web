package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main(){
	//start server

	http.HandleFunc("/",  homeHandler)
	http.HandleFunc("/ascii-art",  asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	    

	

}
//homeHandler
func homeHandler(w http.ResponseWriter, r *http.Request){
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
func asciiArtHandler(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}
		err := r.ParseForm()
		if err != nil{
			http.Error(w, "Bad request", http.StatusBadRequest)
			return

		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" || banner == ""{
			http.Error(w, "Missing input", http.StatusBadRequest)
			return
		}

	result, err := AsciiArt(text, banner)
	if err != nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data:= struct{
		Result string
	}{
		Result: result,
	}
	tmpl.Execute(w, data)
	}

	