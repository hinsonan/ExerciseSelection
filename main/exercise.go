package main

import (
	"html/template"
	"log"
	"net/http"
)

func excerciseSelection(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/exercises.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	//these 2 lines of code serve up the css and js so the template can use them
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", excerciseSelection)
	http.ListenAndServe(":8080", nil)
}
