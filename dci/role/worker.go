package role

import (
	"dci/dci/data"
	"fmt"
)

type Worker struct {
	// Worker同时也是个普通人，因此组合了Human角色
	HumanTrait
	data.WorkCard
}

func (w *Worker) Compose(trait HumanTrait) {
	w.HumanTrait = trait
}

func (w *Worker) Work() {
	fmt.Printf("%+v working\n", w.WorkCard)
	w.CastHuman().Balance++
}

func (w *Worker) OffWork() {
	fmt.Printf("%+v getting off work\n", w.WorkCard)
}
