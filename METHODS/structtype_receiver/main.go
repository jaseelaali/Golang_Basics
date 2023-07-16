/*


METHODS

Go methods are similar to Go function with one difference,
i.e, the method contains a receiver argument in it.
With the help of the receiver argument, the method can access
the properties of the receiver.
Here, the receiver can be of struct type or non-struct type					*/

package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Subject []string
}

func (s  Student) ChangeSubject() {
	s.Subject = append(s.Subject, "Physics")
	fmt.Println(s.Subject)
}
func (s Student) Check() string {
	if s.Age > 12 {
		return s.Name
	}
	return ""
}
func main() {
	s1 := Student{"Anu", 12, []string{"English", "Chemistry", "Maths"}}
	s2 := Student{"Tinu", 13, []string{"Hindi", "Biology", "Maths"}}
	s1.ChangeSubject()
	s2.ChangeSubject()
	fmt.Println(s1.Check())
	fmt.Println(s2.Check(),"have more than 12 year ")
}
