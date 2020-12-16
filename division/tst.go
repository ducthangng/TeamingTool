package division

import (
	"sort"

	"github.com/ducthangng/TeamingTool/db"
)

//SortGenA used to sort according to PA
func SortGenA(gen []db.Student) {
	sort.SliceStable(gen, func(i, j int) bool {
		return gen[i].PA <= gen[j].PA
	})
}

//SortGenO used to sort according to PO
func SortGenO(gen []db.Student) {
	sort.SliceStable(gen, func(i, j int) bool {
		return gen[i].PO <= gen[j].PO
	})
}

//SortTeam used to sort according to number of people in a team
func SortTeam(team []db.Team) {
	sort.SliceStable(team, func(i, j int) bool {
		return len(team[i].List) <= len(team[j].List)
	})
}
