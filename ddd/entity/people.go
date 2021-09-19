package entity

import (
	"dci/ddd/vo"
	"fmt"
	"math/rand"
)

type People struct {
	vo.IdentityCard
	vo.StudentCard
	vo.WorkCard
	vo.Account
}

func NewPeople(name string) *People {
	return &People{
		IdentityCard: vo.IdentityCard{
			Id:   rand.Uint32(),
			Name: name,
		},
		Account: vo.Account{
			Id:      rand.Uint32(),
			Balance: 10,
		},
	}
}

func (p *People) Study() {
	fmt.Printf("Student %+v studying\n", p.StudentCard)
}

func (p *People) Exam() {
	fmt.Printf("Student %+v examing\n", p.StudentCard)
}
func (p *People) Eat() {
	fmt.Printf("%+v eating\n", p.IdentityCard)
	p.Account.Balance--
}

func (p *People) Sleep() {
	fmt.Printf("%+v sleeping\n", p.IdentityCard)
}

func (p *People) PlayGame() {
	fmt.Printf("%+v playing game\n", p.IdentityCard)
}

func (p *People) Work() {
	fmt.Printf("%+v working\n", p.WorkCard)
	p.Account.Balance++
}

func (p *People) OffWork() {
	fmt.Printf("%+v getting off work\n", p.WorkCard)
}

func (p *People) BuyTicket() {
	fmt.Printf("%+v buying a ticket\n", p.IdentityCard)
	p.Account.Balance--
}

func (p *People) Enjoy() {
	fmt.Printf("%+v enjoying scenery\n", p.IdentityCard)
}
