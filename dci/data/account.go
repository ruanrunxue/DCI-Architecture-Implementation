package data

type Account struct {
	Id      uint32
	Balance int
}

type AccountGetter interface {
	GetAccount() *Account
}
