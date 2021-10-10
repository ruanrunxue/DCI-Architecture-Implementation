package service

import (
	"dci/ddd/entity"
	"fmt"
)

type EnjoyerService struct{}

func (e *EnjoyerService) BuyTicket(p *entity.People) {
	fmt.Printf("%+v buying a ticket\n", p.IdentityCard)
	p.Account.Balance--
}

func (e *EnjoyerService) Enjoy(p *entity.People) {
	fmt.Printf("%+v enjoying scenery\n", p.IdentityCard)
}
