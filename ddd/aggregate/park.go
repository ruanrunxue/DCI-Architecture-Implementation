package aggregate

import (
	"dci/ddd/entity"
	"fmt"
)

type Park struct {
	Name     string
	enjoyers []*entity.People
}

func NewPark(name string) *Park {
	park := &Park{}
	park.Name = name
	park.enjoyers = make([]*entity.People, 0)
	return park
}

func (p *Park) Welcome(enjoyer *entity.People) {
	fmt.Printf("%+v come to park %s\n", enjoyer.IdentityCard, p.Name)
	p.enjoyers = append(p.enjoyers, enjoyer)
}

func (p *Park) Run() {
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
