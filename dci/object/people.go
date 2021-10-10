package object

import (
	"dci/dci/data"
	"dci/dci/role"
	"math/rand"
)

type People struct {
	// People对象扮演的角色
	role.Human
	role.Student
	role.Worker
	role.Enjoyer
}

func NewPeople(name string) *People {
	people := &People{
		Human: role.Human{
			IdentityCard: data.IdentityCard{
				Id:   rand.Uint32(),
				Name: name,
			},
			Account: data.Account{
				Id:      rand.Uint32(),
				Balance: 10,
			},
		},
	}
	// 初始化各角色
	people.Student.Compose(people)
	people.Worker.Compose(people)
	people.Enjoyer.Compose(people)
	return people
}

func (p *People) CastHuman() *role.Human {
	return &p.Human
}

func (p *People) CastStudent() *role.Student {
	return &p.Student
}

func (p *People) CastWorker() *role.Worker {
	return &p.Worker
}

func (p *People) CastEnjoyer() *role.Enjoyer {
	return &p.Enjoyer
}
