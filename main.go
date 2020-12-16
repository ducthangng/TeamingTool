package main

import (
	"github.com/ducthangng/TeamingTool/db"
	"github.com/ducthangng/TeamingTool/division"
)

func main() {
	students := []db.Student{
		{ID: 1, Name: "DucThang", Age: 16, PA: 80, PC: 60, PE: 70, PO: 74, PN: 30, Gender: "male"},
		{ID: 2, Name: "DucThang1", Age: 18, PA: 50, PC: 70, PE: 30, PO: 70, PN: 54, Gender: "male"},
		{ID: 3, Name: "DucThang2", Age: 17, PA: 96, PC: 40, PE: 82, PO: 60, PN: 35, Gender: "male"},
		{ID: 4, Name: "DucThang3", Age: 17, PA: 73, PC: 46, PE: 52, PO: 70, PN: 31, Gender: "male"},
		{ID: 5, Name: "DucThang4", Age: 18, PA: 80, PC: 43, PE: 57, PO: 78, PN: 25, Gender: "female"},
		{ID: 6, Name: "DucThang5", Age: 17, PA: 73, PC: 32, PE: 40, PO: 65, PN: 30, Gender: "female"},
		{ID: 7, Name: "DucThang6", Age: 18, PA: 91, PC: 70, PE: 30, PO: 74, PN: 19, Gender: "female"},
		{ID: 8, Name: "DucThang7", Age: 18, PA: 79, PC: 42, PE: 50, PO: 34, PN: 59, Gender: "male"},
		{ID: 9, Name: "DucThang8", Age: 18, PA: 80, PC: 60, PE: 55, PO: 40, PN: 60, Gender: "male"},
		{ID: 10, Name: "DucThang9", Age: 18, PA: 59, PC: 80, PE: 22, PO: 78, PN: 59, Gender: "male"},
		{ID: 11, Name: "DucThang10", Age: 19, PA: 83, PC: 42, PE: 80, PO: 45, PN: 50, Gender: "female"},
		{ID: 12, Name: "DucThang11", Age: 18, PA: 89, PC: 91, PE: 50, PO: 60, PN: 68, Gender: "female"},
		{ID: 13, Name: "DucThang12", Age: 17, PA: 70, PC: 80, PE: 55, PO: 80, PN: 40, Gender: "male"},
		{ID: 14, Name: "DucThang13", Age: 17, PA: 60, PC: 78, PE: 68, PO: 90, PN: 29, Gender: "male"},
		{ID: 15, Name: "DucThang14", Age: 18, PA: 60, PC: 60, PE: 61, PO: 40, PN: 70, Gender: "female"},
		{ID: 16, Name: "DucThang15", Age: 17, PA: 70, PC: 75, PE: 70, PO: 58, PN: 92, Gender: "female"},
	}

	division.TeamDivision(students, 4)
}
