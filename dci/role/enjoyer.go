package role

import (
	"fmt"
)

type Enjoyer struct {
	// Enjoyer同时也是个普通人，因此组合了Human角色
	CastedRoles enjoyerCastedRoles
}

type enjoyerCastedRoles interface {
	HumanRole
}

func (e *Enjoyer) BuyTicket() {
	fmt.Printf("%+v buying a ticket\n", e.CastedRoles.CastHuman().IdentityCard)
	e.CastedRoles.CastHuman().Balance--
}

func (e *Enjoyer) Enjoy() {
	fmt.Printf("%+v enjoying scenery\n", e.CastedRoles.CastHuman().IdentityCard)
}
