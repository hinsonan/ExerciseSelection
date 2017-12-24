package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	data "github.com/andrewhinson/ExerciseSelection/data"
	models "github.com/andrewhinson/ExerciseSelection/models"
)

var allData = data.DataArray

func getNewExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	io.WriteString(w, allData[0].Name)

}

func excerciseSelection(w http.ResponseWriter, r *http.Request) {
	var s models.Workout
	if r.FormValue("bicep") == "on" {
		s = allData[0]
	}

	t, err := template.ParseFiles("templates/exercises.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, s)
	if err != nil {
		log.Print("template executing error: ", err)
	}

}

func main() {
	//these 2 lines of code serve up the css and js so the template can use them
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", excerciseSelection)
	http.HandleFunc("/newExercise", getNewExercise)
	http.ListenAndServe(":8080", nil)
}
