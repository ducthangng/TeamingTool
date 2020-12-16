package division

import (
	"github.com/ducthangng/TeamingTool/db"
)

// TeamDivision takes in the students with the same age, which mean age differentiation is
// maximun 2 units.
func TeamDivision(s []db.Student, indexLimits int) {
	var finalfilter []db.Student
	var teamA []db.Team
	var teamO []db.Team
	var teamX []db.Team
	var ateamX db.Team

	students, redoers := CheckOverallReq(s)

	// Team A division
	male, female := Genderize(students)
	SortGenA(male)
	SortGenA(female)
	limitX, limitY, leftoverA := TypeAReq(&male, &female)
	teamA, conA := Typing(male, female, limitX, limitY)

	// Merge the leftover to the continuous
	conA = append(conA, leftoverA...)

	// Team E division
	male, female = Genderize(conA)
	SortGenO(male)
	SortGenO(female)
	limitX, limitY, leftoverO := TypeOReq(&male, &female)
	teamO, conO := Typing(male, female, limitX, limitY)
	// Distribute to team A and E the right member
	// At this state, we have: team A - team O - Con: the leftover who cannot join team. - The unqualifieds
	// There are various reasons can be listed here for the Con:
	// 1. Many students with same gender -> Cannot be made into team
	// 2. Just some leftover students.
	// How to deal with this: only distribute it into others team.

	conO = append(conO, leftoverO...)
	label := Labelling(conO)

	for i, v := range conO {
		k := 0
		check := false
		for j, v2 := range teamO {
			if (label[i] == "E" || label[i] == "O") && (len(v2.List) < indexLimits && (len(v2.List) != 0)) {
				k = j
				check = true
				break
			}
		}
		if check == true {
			teamO[k].List = append(teamO[k].List, v)
		} else {
			if check == false {
				for j, v2 := range teamA {
					if (label[i] == "A" || label[i] == "C" || label[i] == "N") && (len(v2.List) < indexLimits) && (j != 0) {
						k = j
						check = true
						break
					}
				}
			}
			if check == true {
				teamA[k].List = append(teamA[k].List, v)
			}
		}

		if check == false {
			finalfilter = append(finalfilter, v)
		}
	}

	for _, v := range finalfilter {
		ateamX.List = append(ateamX.List, v)
	}
	teamX = append(teamX, ateamX)

	ShowTeam("Type A teams", teamA)
	ShowTeam("Type O teams", teamO)
	ShowTeam("Type X teams", teamX)
	ShowTeam("Redoers: ", redoers)
}
