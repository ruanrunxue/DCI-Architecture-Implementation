package aggregate

import (
	"dci/ddd/entity"
	"dci/ddd/vo"
	"fmt"
	"math/rand"
)

type School struct {
	Name     string
	students []*entity.People
}

func NewSchool(name string) *School {
	school := &School{}
	school.Name = name
	school.students = make([]*entity.People, 0)
	return school
}

func (s *School) Receive(student *entity.People) {
	student.StudentCard = vo.StudentCard{
		Id:     rand.Uint32(),
		Name:   student.IdentityCard.Name,
		School: s.Name,
	}
	s.students = append(s.students, student)
	fmt.Printf("%s Receive stduent %+v\n", s.Name, student.StudentCard)
}

func (s *School) Run() {
	fmt.Printf("%s start class\n", s.Name)
	for _, student := range s.students {
		student.Study()
	}
	fmt.Println("students start to eating")
	for _, student := range s.students {
		student.Eat()
	}
	fmt.Println("students start to exam")
	for _, student := range s.students {
		student.Exam()
	}
	fmt.Printf("%s finish class\n", s.Name)
}
