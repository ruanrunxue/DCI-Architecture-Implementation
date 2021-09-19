package context

import (
	"dci/dci/role"
	"fmt"
)

type Park struct {
	Name     string
	enjoyers []*role.Enjoyer
}

func NewPark(name string) *Park {
	park := &Park{}
	park.Name = name
	park.enjoyers = make([]*role.Enjoyer, 0)
	return park
}

func (p *Park) Welcome(enjoyer *role.Enjoyer) {
	fmt.Printf("%+v come park %s\n", enjoyer.CastedRoles.CastHuman().IdentityCard, p.Name)
	p.enjoyers = append(p.enjoyers, enjoyer)
}

func (p *Park) Start() {
	fmt.Printf("%s start to sell tickets\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.BuyTicket()
	}
	fmt.Printf("%s start a show\n", p.Name)
	for _, enjoyer := range p.enjoyers {
		enjoyer.Enjoy()
	}
	fmt.Printf("show finish\n")
}
