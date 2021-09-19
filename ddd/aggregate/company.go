package aggregate

import (
	"dci/ddd/entity"
	"dci/ddd/vo"
	"fmt"
	"math/rand"
)

type Company struct {
	Name    string
	workers []*entity.People
}

func NewCompany(name string) *Company {
	company := &Company{}
	company.Name = name
	company.workers = make([]*entity.People, 0)
	return company
}

func (c *Company) Employ(worker *entity.People) {
	worker.WorkCard = vo.WorkCard{
		Id:      rand.Uint32(),
		Name:    worker.IdentityCard.Name,
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
		worker.Eat()
	}
	fmt.Println("worker get off work")
	for _, worker := range c.workers {
		worker.OffWork()
	}
	fmt.Printf("%s finish work\n", c.Name)
}
