package data

type StudentCard struct {
	Id     uint32
	Name   string
	School string
}

type StudentCardGetter interface {
	GetStudentCard() *StudentCard
}
