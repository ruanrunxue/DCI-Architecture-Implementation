package role

import (
	"fmt"
)

type Enjoyer struct {
	// Enjoyer同时也是个普通人，因此组合了Human角色
	HumanTrait
}

func (e *Enjoyer) Compose(trait HumanTrait) {
	e.HumanTrait = trait
}

func (e *Enjoyer) BuyTicket() {
	fmt.Printf("%+v buying a ticket\n", e.CastHuman().IdentityCard)
	e.CastHuman().Balance--
}

func (e *Enjoyer) Enjoy() {
	fmt.Printf("%+v enjoying scenery\n", e.CastHuman().IdentityCard)
}
