package data

type IdentityCard struct {
	Id   uint32
	Name string
}

type IdentityCardGetter interface {
	GetIdentityCard() *IdentityCard
}
