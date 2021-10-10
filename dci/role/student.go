package role

import (
	"dci/dci/data"
	"fmt"
)

type Student struct {
	// Student同时也是个普通人，因此组合了Human角色
	HumanTrait
	data.StudentCard
}

// 学生角色特征
type StudentTrait interface {
	CastStudent() *Student
}

func (s *Student) Compose(trait HumanTrait) {
	s.HumanTrait = trait
}

func (s *Student) Study() {
	fmt.Printf("Student %+v studying\n", s.StudentCard)
}

func (s *Student) Exam() {
	fmt.Printf("Student %+v examing\n", s.StudentCard)
}
