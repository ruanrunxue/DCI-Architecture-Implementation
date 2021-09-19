package aggregate

import (
	"dci/ddd/entity"
	"fmt"
)

type Home struct {
	me *entity.People
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) ComeBack(p *entity.People) {
	fmt.Printf("%+v come back home\n", p.IdentityCard)
	h.me = p
}

func (h *Home) Run() {
	h.me.Eat()
	h.me.PlayGame()
	h.me.Sleep()
}
