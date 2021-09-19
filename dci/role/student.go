package role

import (
	"dci/dci/data"
	"fmt"
)

type Student struct {
	data.StudentCard
	// Student同时也是个普通人，因此组合了Human角色
	CastedRoles studentCastedRoles
}

type studentCastedRoles interface {
	HumanRole
}

func (s *Student) Study() {
	fmt.Printf("Student %+v studying\n", s.StudentCard)
}

func (s *Student) Exam() {
	fmt.Printf("Student %+v examing\n", s.StudentCard)
}
