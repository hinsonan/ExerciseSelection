package data

import models "github.com/andrewhinson/ExerciseSelection/models"

//this will be used to store the exercise data statically
var DataArray = []models.Workout{
	models.Workout{"Bicep Curl", "Isolation exercise that heavily impacts the bicep", "../img/bicep.jpg"},
	models.Workout{"Hammer Curl", "Isolation exercise that heavily impacts the bicep", "../img/hammercurl.jpg"},
	models.Workout{"Squat", "Works the back, hamstrings, quads, glutes, and core", "../img/squat.jpg"},
	models.Workout{"Bench", "Works the triceps, back, shoulders, core, biceps, and chest", "../img/bench.jpg"},
	models.Workout{"Deadlift", "Works the back, quad, bicep, forearm, glutes, hamstrings, and core", "../img/deadlift.jpg"},
	models.Workout{"Skull Crusher", "Isolation work to target the tricep", "../img/skullcrusher.jpg"},
	models.Workout{"Goodmorning", "Works the back and hamstrings", "../img/goodmorning.jpg"},
	models.Workout{"Front Squat", "Works the quads, back, and core", "../img/frontsquat.jpg"},
	models.Workout{"Over Head Press", "Works the shoulders, triceps, and core", "../img/ohp.jpg"},
}
