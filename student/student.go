package main

//interfaces learnning by ehsan.jahangard
import (
	"errors"
	"fmt"
)

// set func to struct (not use interface)
type Car struct {
	Color string
	Name  string
}

func (c Car) GetColor() string {

	//fmt.Println("my car is name:", c.Name)
	//fmt.Println("my car is color:", c.Color)
	return "my car is name:" + c.Name + "--" + "my car is color:" + c.Color
}

// step A
type IStudentProp interface {
	GetFullName() string
	GetCar() (string, error)
}

// step B
// model
type Student struct {
	Name   string
	Family string
}

// Step C (implimentaition) Link to Model
func (s Student) GetFullName() string {
	return s.Name + " " + s.Family
}
func main() {
	//create
	studentModel := &Student{Name: "JJ", Family: "Ehsan"}
	//use
	res := studentModel.GetFullName()
	fmt.Printf(res)
	fmt.Printf("\n**********************\t****\t\r*****************\n")
	c1 := &Car{Color: "red", Name: "Benz"}
	fmt.Println(c1.GetColor())
}

func (s Student) GetCar() (string, error) {
	if s.Name == "" {
		return "", errors.New("erorrrrr")
	}
	return "", fmt.Errorf("")
}
