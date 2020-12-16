package division

import (
	"github.com/ducthangng/TeamingTool/db"
)

//Labelling is used for label the student's most prominent type
func Labelling(s []db.Student) []string {
	var label []string
	for _, v := range s {
		if (v.PA >= v.PC) && (v.PA >= v.PE) && (v.PA >= v.PO) && (v.PA >= v.PN) {
			label = append(label, "A")
		} else if (v.PO >= v.PC) && (v.PO >= v.PE) && (v.PO >= v.PA) && (v.PO >= v.PN) {
			label = append(label, "O")
		} else if (v.PE >= v.PC) && (v.PE >= v.PA) && (v.PE >= v.PO) && (v.PE >= v.PN) {
			label = append(label, "E")
		} else if (v.PC >= v.PA) && (v.PC >= v.PE) && (v.PC >= v.PO) && (v.PC >= v.PN) {
			label = append(label, "C")
		} else if (v.PN >= v.PC) && (v.PN >= v.PE) && (v.PN >= v.PO) && (v.PN >= v.PN) {
			label = append(label, "N")
		}
	}
	return label
}
