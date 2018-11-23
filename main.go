package main

import (
	data "ExerciseSelection/data"
	models "ExerciseSelection/models"
	"html/template"
	"io"
	"log"
	"net/http"
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

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "bicep" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("quad") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "quad" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("back") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "back" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("hamstring") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "hamstring" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("chest") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "chest" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("tricep") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "tricep" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	if r.FormValue("shoulder") == "on" {

		//filters the correct data based on user input
		for _, element := range allData {

			if element.PrimaryMuscle == "shoulder" {
				//we have to append the data so we dont go out of bounds
				workoutData = append(workoutData, element)
			}
		}

	}

	//sets the template up and executes it with the data
	t, err := template.ParseFiles("templates/index.html")
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
