package division

import "github.com/ducthangng/TeamingTool/db"

// TypeAReq is used for pointing out safely the limit of A. This function return 0 <= x < len(gender) if the type has a limit. Else it would return len(gender).
func TypeAReq(male *[]db.Student, female *[]db.Student) (int, int, []db.Student) {
	var x, y int
	var t, f bool
	var leftover []db.Student
	t = false
	f = false
	for i, v := range *male {
		if v.PA >= 60 {
			x = i
			t = true
			break
		}
		leftover = append(leftover, v)
	}
	for i, v := range *female {
		if v.PA >= 60 {
			y = i
			f = true
			break
		}
		leftover = append(leftover, v)
	}
	if (t == true) && (f == true) {
		return x, y, leftover
	}
	if (t == true) && (f == false) {
		return x, len(*female), leftover
	}
	if (t == false) && (f == true) {
		return len(*male), len(*female), leftover
	}
	return len(*male), len(*female), leftover
}

// TypeOReq is used for pointing out safely the limit of O. This function return 0 <= x < len(gender) if the type has a limit. Else it would return len(gender).
func TypeOReq(male *[]db.Student, female *[]db.Student) (int, int, []db.Student) {
	var x, y int
	var t, f bool
	var leftover []db.Student
	t = false
	f = false
	for i, v := range *male {
		if v.PO >= 60 {
			x = i
			t = true
			break
		}
		leftover = append(leftover, v)
	}
	for i, v := range *female {
		if v.PO >= 60 {
			y = i
			f = true
			break
		}
		leftover = append(leftover, v)
	}
	if (t == true) && (f == true) {
		return x, y, leftover
	}
	if (t == true) && (f == false) {
		return x, len(*female), leftover
	}
	if (t == false) && (f == true) {
		return len(*male), len(*female), leftover
	}
	return len(*male), len(*female), leftover
}

// CheckOverallReq is used for getting the qualified students. Returns 2 variables: the students and the Redoers.
func CheckOverallReq(s []db.Student) ([]db.Student, []db.Team) {
	var e []db.Student
	var rb []db.Team
	var n db.Team

	for _, v := range s {
		if (v.PA < 60) && (v.PO < 60) {
			n.List = append(n.List, v)
		} else {
			e = append(e, v)
		}
	}

	rb = append(rb, n)

	return e, rb
}
