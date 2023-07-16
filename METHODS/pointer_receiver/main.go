package main

import "fmt"

type Student struct {
	Name   string
	Branch string
}

func (s *Student) ChangeBranch(branch string) {
	s.Branch = branch

}
func main() {
	stud := Student{"Anu", "cse"}
	fmt.Printf("before\n")
	fmt.Println(stud)
	//p := &stud
	stud.ChangeBranch("ece")
	fmt.Printf("after\n")
	fmt.Println(stud)
}
