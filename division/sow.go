package division

import (
	"fmt"

	"github.com/ducthangng/TeamingTool/db"
)

//ShowTeam is a debug function
func ShowTeam(field string, t []db.Team) {
	fmt.Println(field)
	for i, v := range t {
		fmt.Printf("team: %v \n", i+1)
		for _, v2 := range v.List {
			fmt.Println(v2)
		}
		fmt.Println(" ")
	}
}
