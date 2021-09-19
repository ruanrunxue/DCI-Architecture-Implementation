package service

import (
	"dci/ddd/entity"
	"fmt"
)

type StudentService struct{}

func (s *StudentService) Study(p *entity.People) {
	fmt.Printf("Student %+v studying\n", p.StudentCard)
}

func (s *StudentService) Exam(p *entity.People) {
	fmt.Printf("Student %+v examing\n", p.StudentCard)
}
