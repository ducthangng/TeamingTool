package division

import "github.com/ducthangng/TeamingTool/db"

//Typing is used for type-generating students.
func Typing(male []db.Student, female []db.Student, limitX int, limitY int) ([]db.Team, []db.Student) {
	var code int
	// code = 1 : 1 male 1 female
	// code = 2 : out of female
	// code = 3 : out of male
	// code = 4 : perfect
	// code = 5 : 1 female and a bunch of male
	var final []db.Team
	i, j := limitX, limitY
	h, k := len(male)-1, len(female)-1
	for {
		var team db.Team

		// 1st condition: 2 males 2 females
		if h-1 >= i {
			if k-1 >= j {
				team.List = append(team.List, male[h], male[h-1], female[k], female[k-1])
				k = k - 2
				h = h - 2
			}
		} else if h == i { // 2nd condition: 1 male and 2 females
			if k-1 >= j {
				team.List = append(team.List, male[h], female[k], female[k-1])
				k = k - 2
				h = h - 1
			}
		}
		final = append(final, team)
		if (k == j) && (i == h) {
			code = 1
			break
		}
		if (h >= i) && (j > k) {
			code = 2
			break
		}
		if (h < i) && (j <= k) {
			code = 3
			break
		}
		if (h < i) && (k < j) {
			code = 4
			break
		}
		if (j == k) && (h > i) {
			code = 5
			break
		}
	}
	var leftover []db.Student
	// leftover is the student which cannot be divided
	for {
		if i <= h {
			leftover = append(leftover, male[h])
			h = h - 1
		}
		if j <= k {
			leftover = append(leftover, female[k])
			k = k - 1
		}
		if h < i && k < j {
			break
		}
	}
	return final, leftover

	// the func ends here.
	// code = 1 : 1 male 1 female
	if code == 1 {
		for k := 0; k <= len(final)-1; k++ {
			if len(final[k].List) == 3 {
				if j == len(female)-1 {
					final[k].List = append(final[k].List, female[j])
					j = j + 1
				} else {
					final[k].List = append(final[k].List, male[i])
					i = i + 1
				}
			}
		}
		if i == len(male)-1 {
			final[len(final)-2].List = append(final[len(final)-2].List, male[i])
		}
		if j == len(female)-1 {
			final[len(final)-1].List = append(final[len(final)-1].List, female[j])
		}
	} else if code == 2 {
		// code = 2 : out of female
		listteam, lef := GroupWithSameGender(male, i)
		for _, v := range listteam {
			final = append(final, v)
		}

		for i, v := range lef {
			final[i].List = append(final[i].List, v)
		}
	} else if code == 3 {
		// code = 3 : out of male
		listteam, lef := GroupWithSameGender(female, j)

		for _, v := range listteam {
			final = append(final, v)
		}

		for i, v := range lef {
			final[len(final)-1-i].List = append(final[len(final)-1-i].List, v)
		}
	} else if code == 5 {
		// code = 5 : 1 female and a bunch of male
		listteam, lef := GroupWithSameGender(male, i)
		for _, v := range listteam {
			final = append(final, v)
		}

		for i, v := range lef {
			final[len(final)-1-i].List = append(final[len(final)-1-i].List, v)
		}

		final[len(final)-1].List = append(final[len(final)-1].List, female[j])
	}
	return final, leftover
}
