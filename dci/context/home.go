package context

import (
	"dci/dci/role"
	"fmt"
)

type Home struct {
	me *role.Human
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) ComeBack(human *role.Human) {
	fmt.Printf("%+v come back home\n", human.IdentityCard)
	h.me = human
}

func (h *Home) Start() {
	h.me.Eat()
	h.me.PlayGame()
	h.me.Sleep()
}
