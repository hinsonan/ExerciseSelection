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

//this funtion will execute at the root and will be used to display the exercises
func excerciseSelection(w http.ResponseWriter, r *http.Request) {
	//be used to send the filtered data to template
	var workoutData []models.Workout
	if r.FormValue("bicep") == "on" {
		//filters the correct data based on usr input
		var filteredData = []models.Workout{
			allData[0],
		}
		//fills the data sent to template with the correct filtered data
		workoutData = filteredData
	}
	//sets the template up and executes it with the data
	t, err := template.ParseFiles("templates/exercises.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, workoutData)
	if err != nil {
		log.Print("template executing error: ", err)
	}

}

func main() {
	//these 3 lines of code serve up the css, img and js so the template can use them
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", excerciseSelection)
	http.HandleFunc("/newExercise", getNewExercise)
	http.ListenAndServe(":8080", nil)
}
