package students

import "fmt"

type Student struct {
	name       string
	age        int
	schoolName string
}

func (s *Student) Init(n string, a int, sn string) {
	s.name = n
	s.age = a
	s.schoolName = sn
}

func (s Student) GetInfo() {
	fmt.Println(s.name, s.age, s.schoolName)
}

func (s *Student) ChangeAge(newAge int) {
	s.age = newAge
}

func (s *Student) ChangeSchool(newSchool string) {
	s.schoolName = newSchool
}
