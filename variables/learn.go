package variables

import (
	"fmt"
	"reflect"
	"strconv"
)

var fulllname string
var email = "naphat.d@gmail.com"
var c, python bool = true, false
const vat int = 7


func Leran() {
	fulllname = "naphat"
	fmt.Println(fulllname)
	fmt.Printf("Hello %s email %s \n", fulllname, email)
	fmt.Println(python, c)

	companyName := "RFD"
	isShow := true
	age := 10
	fmt.Println(companyName, isShow, age)
	fmt.Printf("%T\n", isShow)
	fmt.Printf("%v\n",reflect.TypeOf(age))
	
	f := float64(age)
	fmt.Printf("%T \n",f)
	fmt.Printf("%T \n",vat)
	s := strconv.Itoa(vat) //convert type
	fmt.Println(s)
	fmt.Printf("%T \n",s)
	fmt.Println("vat is " + s)
}
