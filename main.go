package main

import (
	"fmt"

	"example.com/go-demo1/mascot"
)

type staff struct {
	name    string
	address string
}

func main() {
	fmt.Println(" Hello")
	fmt.Println(mascot.Add(1, 2))

	// All about Pointers
	x := 12
	p := &x // pointer p is pointing to the memory address of x
	fmt.Println(p)
	fmt.Println(*p) // we are dereferencing the content in memory address p

	// Structs --another very important topic

	type students struct {
		id   int
		name string
	}

	v := students{}
	v.id = 123
	v.name = " Mahesh"

	fmt.Println("id of a student is ", v.id)
	fmt.Println("name of the student is ", v.name)

	// Arrays

	var a [15]int

	for i := 0; i < 10; i++ {

		a[i] = i

		fmt.Println(a[i])
	}

	// Slices
	var s []int = a[0:4]
	fmt.Println(s)

	// another way of defining array

	names := [4]string{
		"Mahesh",
		"John",
		"George",
		"Ringo",
	}

	// slicers are like references to the arrays
	fmt.Println(names)
	b := names[0:4] // defining the slicer
	fmt.Println(b)
	b[0] = " MaheshDOn" // changing the value of slicer changes the first arrays used
	fmt.Println(b)
	fmt.Println(names)

	// Slice Literals
	q := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(q)

	// slice length and capaciy

	m := []int{2, 3, 5, 7, 9}
	fmt.Println(len(m), cap(m), m)

	// the length changes but capacity remains same- original size of array
	m = m[:3]
	fmt.Println(len(m), cap(m))

	// nil slices
	n := []int{}

	fmt.Println(len(n), cap(n), n)

	// creating slice with make
	l := make([]int, 5)    // both length and capacity is 5
	o := make([]int, 2, 5) // length is 2 and capacity is 5

	fmt.Println(len(l), cap(l), len(o), cap(o))

	// Double dimensional arrays
	mat := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 13},
	}

	fmt.Println(mat[0][0])
	fmt.Println(mat[1][1])
	fmt.Println(mat[2][0])

	// Slices can be used as ArrayList
	matri := []string{}
	matri = append(matri, "Mahesh")
	fmt.Println(matri)
	matri = append(matri, " Hari")
	fmt.Println(matri[0])

	// range can extract elements from slices or array one by one
	// i will auto increment itself

	arrRange := []int{30, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range arrRange {
		fmt.Printf("i is %d v is %d equals %d \n", i, v, (i * v))
	}

	// can escape the index or value by assigning _

	// Meaning we do not even need index , can replace with _ if we are not worried about index
	for _, q := range arrRange {
		fmt.Println(q)
	}

	// maps for key values pair

	type emp struct {
		id   int
		name string
	}

	var mapp = make(map[string]emp)
	mapp["ID Name"] = emp{101, "Mahesh"}
	fmt.Println(mapp["ID Name"])

	// map literals

	type products struct {
		cost  int
		sales int
	}

	// maps are like struct literals but required keys and
	//records are displayed in a sorted way, automatically sorted
	var detail = map[string]products{
		"Nissan_S":      products{100, 200},
		"Toyota_S":      products{20, 50},
		"Lamborghini_s": products{1000, 5000},
	}
	fmt.Println(detail)

	// we donot need to define struct inside map if its type is defined while defining map
	var detail2 = map[string]products{
		"Nissan_S":      {100, 200},
		"Toyota_S":      {20, 50},
		"Lamborghini_s": {1000, 5000},
	}
	fmt.Println(detail2)

	// mapping -- Mutations

	mapps := make(map[string]int)
	mapps["Answer"] = 4
	fmt.Println(" the value of  :", mapps["Answer"])

	mapps["Answer"] = 23
	fmt.Println(" the value is ", mapps["Answer"])

	// when we delete , the value will be zero 0
	delete(mapps, "Answer")
	fmt.Println(" the value is now", mapps["Answer"])

	// searching we have the key present in the map
	//mapps["Answer"] = 29
	// w represents the value in key and ok returns if the key exists in map
	w, ok := mapps["Answer"]
	fmt.Println(" The value :", w, " Present ?", ok)

	// Functiion values

	hypot := func(x, y int) int {
		return x + y
	}

	// comp function will call hypot and get result from hypo and return that result here as sum
	comp := func(fn func(int, int) int) int {
		return fn(3, 4)
	}

	fmt.Println(hypot(1, 2))
	fmt.Println(comp(hypot))

	// function closures

	adder := func() func(int) int {
		sum := 0
		return func(x int) int { // the adder function is called in pos and pos is now function func(x int)int
			sum = sum + x // we then call pos function with 1 as param, and got sum
			return sum    // adder returns closure function and each closure is bound to its own sum
		}
	}

	pos := adder()

	fmt.Println(pos(1 * -2))
	fmt.Println(5)

	// Go does not have classes
	// Methods are functions- with method only having reciever arguments

	// WE CAN DECLARE METHODS WITH POINTER RECIEVER TOO
	// FOR EXAMPLE-- func(v *Vertex)abs() int{}
	// FUNCTION WITH POINTER ARGUMENT MUST TAKE A POINTER
	//

	staf := staff{"Mahesh --", "--363 Main Street Everett MA 02149"}
	//	fmt.Println(staf.getEmp())
	fmt.Println((getEmp(staf))) // CALLING getEmp FUNCTION
	// could not implement interface at the moment
	// var ab Abser
	// myemp:=getEmp(staf)
	// ab=myemp

}

// ***********OUTSIDE THE MAIN FUNCTION******************8

// func (v staff) getEmp() string {
// 	return (v.name + v.address)
// }
type Abser interface {
	abs() int
}

func getEmp(v staff) string {
	return (v.name + v.address)
}
