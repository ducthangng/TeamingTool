package division

import (
	"fmt"

	"github.com/ducthangng/TeamingTool/db"
)

//DistributeIntoGroup is currently implemented in team.go
func DistributeIntoGroup(ta []db.Team, te []db.Team, con []db.Student) {
	label := Labelling(con)
	fmt.Println(label)
	for i, v := range con {
		fmt.Printf("%v is labeled %v: \n", v, label[i])
		check := false
		for _, v2 := range te {
			if (label[i] == "E" || label[i] == "O") && len(v2.List) < 5 {
				v2.List = append(v2.List, v)
				fmt.Println(v2.List)
				check = true
			}
		}
		if check == false {
			for _, v2 := range ta {
				if (label[i] == "A" || label[i] == "C" || label[i] == "N") && len(v2.List) < 5 {
					v2.List = append(v2.List, con[i])
					fmt.Println(v2.List)
					check = true
				}
			}
		}
		if check == false {
			fmt.Printf("Cannot distribute this person: %v\n", v)
		}
	}
}
