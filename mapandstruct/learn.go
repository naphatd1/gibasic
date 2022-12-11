package mapandstruct

import "fmt"

func Learn() {
	// Maps (key, value)

	m := map[string]int{"token": 100, "access": 20}
	fmt.Println(m)
	fmt.Println(m["access"])
	m["access"] = 200
	fmt.Println(m["access"])

	// loop
	for key, v := range m {
		fmt.Printf("%v %v \n", key, v)
	}

	//delete map
	delete(m, "access")
	for key, v := range m {
		fmt.Printf("%v %v \n", key, v)
	}

	// add map
	m["golang"] = 400
	fmt.Println(m)
	for key, v := range m {
		fmt.Printf("%v %v \n", key, v)
	}

	// make map
	var m2 = make(map[string]int)
	m2["a"] = 10
	m2["b"] = 20
	fmt.Println(m2["b"])

	// struct

}
