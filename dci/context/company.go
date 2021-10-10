package context

import (
	"dci/dci/data"
	"dci/dci/role"
	"fmt"
	"math/rand"
)

type Company struct {
	Name    string
	workers []*role.Worker
}

func NewCompany(name string) *Company {
	company := &Company{}
	company.Name = name
	company.workers = make([]*role.Worker, 0)
	return company
}

func (c *Company) Employ(worker *role.Worker) {
	worker.WorkCard = data.WorkCard{
		Id:      rand.Uint32(),
		Name:    worker.CastHuman().IdentityCard.Name,
		Company: c.Name,
	}
	c.workers = append(c.workers, worker)
	fmt.Printf("%s Employ worker %s\n", c.Name, worker.WorkCard.Name)
}

func (c *Company) Run() {
	fmt.Printf("%s start work\n", c.Name)
	for _, worker := range c.workers {
		worker.Work()
	}
	fmt.Println("worker start to eating")
	for _, worker := range c.workers {
		worker.CastHuman().Eat()
	}
	fmt.Println("worker get off work")
	for _, worker := range c.workers {
		worker.OffWork()
	}
	fmt.Printf("%s finish work\n", c.Name)
}
