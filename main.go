package main

import (
	"fmt"
	"os"
	person "sipme/Models/Person"
	calcService "sipme/SharedMethod"
)

var name string = "ehsan"

const (
	ip = "ehsan"
)

func main() {
	calcService.InitPack()
	calcService.CALCCARTING()
	calcService.CALCCARTING()
	file, err := os.Create("ehsanGoLangText.txt")
	if err != nil {
		fmt.Println("file error create")
	}
	file.Write([]byte("ehsan write by go lang"))
	fileinfo, err := os.Stat("ehsanGoLangText.txt")
	fmt.Println("fileinfo : ", fileinfo)
	file.Close()

	//array
	// var a [6]int
	// a = [6]int{1, 2, 3, 44, 5, 6}
	// for index, value := range a {

	// 	fmt.Println(index, value)
	// }
	//slice (dynamic array)
	var b []int
	b = append(b, 5)
	fmt.Println(cap(b))
	fmt.Println(len(b))
	var p1 person.Person
	p1.Age = 10
	p1.FirstName = "ehsan"
	p1.LastName = "jj"
	p1.ID = 1
	//var p2 = person.Person{ID: 1, FirstName: "ehsan", LastName: "jjj", Age: 10}
	//var v int = 10
	//fmt.Println("hey", rand.Intn(10))
	//fmt.Println("hey22222", calcPie(v, 3))
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Tracert", i)
	// }
	// for 1 == 1 {
	// 	v += 1
	// 	if v == 50 {
	// 		fmt.Println("break", v)
	// 		break
	// 	}

	// }
	var m map[string]person.Person
	m = make(map[string]person.Person)
	m["a"] = person.Person{ID: 1, FirstName: "new name student", LastName: "family"}
	if m == nil {
		fmt.Println("nill data")
	}
	fmt.Println("map collection", m["a"])
	delete(m, "a")

	type Student struct {
		ID   uint
		Name string
	}
	// make a new instance class
	student1 := new(Student)
	student1.ID = 1
	student1.Name = "ehsan"
	println(student1)
	println(student1.Name)

	// make a new instance class with init value(best practice)
	student2 := &Student{
		ID:   1,
		Name: "ali"}
	println(student2)
	println(student2.Name)

}
func calcPie(x int, y int) int {
	return x * y
}
