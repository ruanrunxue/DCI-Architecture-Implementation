package context

import (
	"dci/dci/data"
	"dci/dci/role"
	"fmt"
	"math/rand"
)

type School struct {
	Name     string
	students []*role.Student
}

func NewSchool(name string) *School {
	school := &School{}
	school.Name = name
	school.students = make([]*role.Student, 0)
	return school
}

func (s *School) Receive(student *role.Student) {
	student.StudentCard = data.StudentCard{
		Id:     rand.Uint32(),
		Name:   student.CastHuman().IdentityCard.Name,
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
		student.CastHuman().Eat()
	}
	fmt.Println("students start to exam")
	for _, student := range s.students {
		student.Exam()
	}
	fmt.Printf("%s finish class\n", s.Name)
}
