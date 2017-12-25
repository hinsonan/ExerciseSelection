package data

import models "github.com/andrewhinson/ExerciseSelection/models"

//this will be used to store the exercise data statically
var DataArray = []models.Workout{
	models.Workout{"Bicep Curl", "Isolation exercise that heavily impacts the bicep", "../img/bicep.jpg"},
	models.Workout{"Hammer Curl", "Isolation exercise that heavily impacts the bicep", "../img/bicep.jpg"},
	models.Workout{"Squat", "Works the back, hamstrings, quads, glutes, and core", "../img/bicep.jpg"},
	models.Workout{"Bench", "Works the triceps, back, shoulders, core, biceps, and chest", "../img/bicep.jpg"},
	models.Workout{"Deadlift", "Works the back, quad, bicep, forearm, glutes, hamstrings, and core", "../img/bicep.jpg"},
	models.Workout{"Tricep Extension", "Isolation work to target the tricep", "../img/bicep.jpg"},
	models.Workout{"Goodmorning", "Works the back and hamstrings", "../img/bicep.jpg"},
	models.Workout{"Front Squat", "Works the quads, back, and core", "../img/bicep.jpg"},
	models.Workout{"Over Head Press", "Works the shoulders, triceps, and core", "../img/bicep.jpg"},
}
