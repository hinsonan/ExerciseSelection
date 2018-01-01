package data

import models "github.com/andrewhinson/ExerciseSelection/models"

//this is an enum in go its made by using constants
const (
	bicep     = "bicep"
	tricep    = "tricep"
	quad      = "quad"
	back      = "back"
	hamstring = "hamstring"
	core      = "core"
	shoulders = "shoulder"
	glutes    = "glutes"
	chest     = "chest"
)

//this will be used to store the exercise data statically
var DataArray = []models.Workout{
	models.Workout{"Bicep Curl", "Isolation exercise that heavily impacts the bicep", "../img/bicep.jpg", bicep, ""},
	models.Workout{"Hammer Curl", "Isolation exercise that heavily impacts the bicep", "../img/hammercurl.jpg", bicep, ""},
	models.Workout{"Squat", "Works the back, hamstrings, quads, glutes, and core", "../img/squat.jpg", quad, glutes},
	models.Workout{"Bench", "Works the triceps, back, shoulders, core, biceps, and chest", "../img/bench.jpg", chest, tricep},
	models.Workout{"Deadlift", "Works the back, quad, bicep, forearm, glutes, hamstrings, and core", "../img/deadlift.jpg", back, hamstring},
	models.Workout{"Skull Crusher", "Isolation work to target the tricep", "../img/skullcrusher.jpg", tricep, ""},
	models.Workout{"Goodmorning", "Works the back and hamstrings", "../img/goodmorning.png", hamstring, back},
	models.Workout{"Front Squat", "Works the quads, back, and core", "../img/frontsquat.jpg", quad, core},
	models.Workout{"Over Head Press", "Works the shoulders, triceps, and core", "../img/ohp.png", shoulders, tricep},
}
