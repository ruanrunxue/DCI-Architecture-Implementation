package role

import (
	"dci/dci/data"
	"fmt"
)

type Worker struct {
	data.WorkCard
	// Worker同时也是个普通人，因此组合了Human角色
	CastedRoles workerCastedRoles
}

type workerCastedRoles interface {
	HumanRole
}

func (w *Worker) Work() {
	fmt.Printf("%+v working\n", w.WorkCard)
	w.CastedRoles.CastHuman().Balance++
}

func (w *Worker) OffWork() {
	fmt.Printf("%+v getting off work\n", w.WorkCard)
}
