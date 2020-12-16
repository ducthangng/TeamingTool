package db

type Student struct {
	ID     int
	Name   string
	Gender string
	Age    int
	PA     int
	PC     int
	PE     int
	PO     int
	PN     int
}

type Team struct {
	List []Student
}
