package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Hello World")
	fmt.Print("Hiiiii!")

	var name string = "fola"
	age := 15
	fmt.Printf("Hello %s", name)
	fmt.Printf("age is %d\n", age)

	fmt.Println()

	number := 50

	if number < 50 {

		fmt.Println(false)

	} else {

		fmt.Println(true)
	}

	validate()

	fmt.Println(validated("Babs"))

	fmt.Println(validates())

	newName := ""
	fmt.Println("Enter name:")
	fmt.Scan(&newName)
	fmt.Println("Hello ", newName)

	fmt.Println(len(newName))

	//for count := 0; count <= 10; count++ {
	//	fmt.Println(count)
	//}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	he := 8
	result := strconv.Itoa(he)
	fmt.Println(result, "2", "4")

}

func validate() {

}

func validated(name string) string {

	if name == "" {

		fmt.Println("Wrong input")
	}

	return name
}

func validates() (int, string, float64) {

	age := 18
	name := "Lawal"

	grade := 12.949

	return age, name, grade

}
