package pointers

import (
	"fmt"
)

func Learn() {
	x := 10

	fmt.Println(x)
	fmt.Println(&x) // & is get memory address

	println("______")
	y := x
	fmt.Println(y)
	fmt.Println(&y)

	println("______")
	p := &x         // pointer p
	fmt.Println(p)  // read data p pointer (dereference)
	fmt.Println(*p) // read data p pointer (dereference)

	println("______")
	*p = 20
	fmt.Println(*p)
	fmt.Println(x)

	println("______")
	ss := student{name: "JJ"}
	fmt.Println(ss.name)
	setName(&ss)
	fmt.Println(ss.name)

}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Ben"
}

// & ops get address
// * get value วางหน้าตัวแปร
// * type get value type วางหน้า type