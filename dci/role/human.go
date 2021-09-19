package role

import (
	"dci/dci/data"
	"fmt"
)

type Human struct {
	data.IdentityCard
	data.Account
}

type HumanRole interface {
	CastHuman() *Human
}

func (h *Human) Eat() {
	fmt.Printf("%+v eating\n", h.IdentityCard)
	h.Account.Balance--
}

func (h *Human) Sleep() {
	fmt.Printf("%+v sleeping\n", h.IdentityCard)
}

func (h *Human) PlayGame() {
	fmt.Printf("%+v playing game\n", h.IdentityCard)
}
