package workoutmodel

type workout struct {
	Name        string
	Description string
}

var dataArray = []workout{
	workout{"bicep", "Works the arm"},
	workout{"squat", "Works the quads"},
}
