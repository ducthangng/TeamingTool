package division

import (
	"fmt"
	"sort"

	"github.com/ducthangng/TeamingTool/db"
)

//AgeDiv is not useable, as age is not considered.
func AgeDiv(s []db.Student) ([]db.Student, []int) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Age < s[j].Age
	})

	var result []db.Student
	var trace []int

	null := db.Student{Name: "NULL"}

	result = append(result, s[0])
	trace = append(trace, 0)

	for i := range s {
		if i == len(s)-1 {
			break
		}
		fmt.Println(s[i+1].Age, " ", s[trace[i]], " ", trace[i])
		if s[i+1].Age-result[trace[i]].Age > 2 {
			result = append(result, null, s[i+1])
			trace = append(trace, 0)
			trace = append(trace, len(result)-1)
		} else {
			result = append(result, s[i+1])
			trace = append(trace, trace[i])
		}
	}

	return result, trace
}
