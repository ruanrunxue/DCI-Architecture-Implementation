package data

type WorkCard struct {
	Id      uint32
	Name    string
	Company string
}

type WorkCardGetter interface {
	GetWorkCard() *WorkCard
}
