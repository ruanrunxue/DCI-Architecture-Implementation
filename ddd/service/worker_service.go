package service

import (
	"dci/ddd/entity"
	"fmt"
)

type WorkerService struct{}

func (w *WorkerService) Work(p *entity.People) {
	fmt.Printf("%+v working\n", p.WorkCard)
	p.Account.Balance++
}

func (w *WorkerService) OffWOrk(p *entity.People) {
	fmt.Printf("%+v getting off work\n", p.WorkCard)
}
