package service

import (
	"dci/ddd/entity"
	"fmt"
)

type HumanService struct{}

func (h *HumanService) Eat(p *entity.People) {
	fmt.Printf("%+v eating\n", p.IdentityCard)
	p.Account.Balance--
}

func (h *HumanService) Sleep(p *entity.People) {
	fmt.Printf("%+v sleeping\n", p.IdentityCard)
}

func (h *HumanService) PlayGame(p *entity.People) {
	fmt.Printf("%+v playing game\n", p.IdentityCard)
}
