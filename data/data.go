package data

import models "github.com/andrewhinson/ExerciseSelection/models"

//this will be used to store the exercise data statically
var DataArray = []models.Workout{
	models.Workout{"Bicep Curl", "Isolation exercise that heavily impacts the bicep"},
	models.Workout{"Squat", "Works the back, hamstrings, quads, glutes, and core"},
	models.Workout{"Bench", "Works the triceps, back, shoulders, core, biceps, and chest"},
	models.Workout{"Deadlift", "Works the back, quad, bicep, forearm, glutes, hamstrings, and core"},
	models.Workout{"Tricep Extension", "Isolation work to target the tricep"},
	models.Workout{"Goodmorning", "Works the back and hamstrings"},
	models.Workout{"Front Squat", "Works the quads, back, and core"},
	models.Workout{"Over Head Press", "Works the shoulders, triceps, and core"},
}
