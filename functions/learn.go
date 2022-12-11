package functions

import (
	"fmt"
	"strings"
)

// var Fullname string

func add(x, y int) int {
	return x + y
}

func converter(title, email string) (string, string) {
	s1 := strings.ToUpper(title)
	return s1, email
}

func sum(numbers ...int) {
	fmt.Println(numbers)
	total := 0
	for _, v := range numbers {
		total += v
	}
	fmt.Println(total)
}

func Learn() {
	fmt.Println(add(10, 5))
	fmt.Println(converter("Hello", "hello.gmail.com"))
	sum(10, 20, 30, 40)

}
