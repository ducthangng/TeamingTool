package division

import (
	"github.com/ducthangng/TeamingTool/db"
)

//Genderize return the classification of male and female students.
func Genderize(s []db.Student) ([]db.Student, []db.Student) {
	var male []db.Student
	var female []db.Student

	for _, v := range s {
		if v.Gender == "male" {
			male = append(male, v)
		} else {
			female = append(female, v)
		}
	}

	return male, female
}

//GroupWithSameGender is currently not supported.
func GroupWithSameGender(gender []db.Student, index int) ([]db.Team, []db.Student) {
	var result []db.Team
	for {
		if len(gender)-index >= 4 {
			var Team db.Team
			Team.List = append(Team.List, gender[index], gender[index+1], gender[index+2], gender[index+3])
			index = index + 4
			result = append(result, Team)
		}
	}
	/*
		for {
			if len(gender)-index >= 3 {
				var Team db.Team
				Team.List = append(Team.List, gender[index], gender[index+1], gender[index+2])
				index = index + 3
				result = append(result, Team)
			}
		}*/
}
