package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {

	score := 10
	if score == 10 {
		fmt.Println("very good")
	} else {
		fmt.Println("try agin")
	}
	os := runtime.GOOS
	fmt.Println("Go run on")
	switch os {
	case "darwin":
		fmt.Println("MAC")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s \n", os)
	}

	b, err := ioutil.ReadFile("file.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	content := string(b)
	fmt.Println(content)

	for i := 10; i < 10; i++ {
		fmt.Println(i)
	}
	for j := 5; j >= 0; j-- {
		if j == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(j)
		time.Sleep(1 * time.Second)
	}
}
