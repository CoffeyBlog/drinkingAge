package main

import "fmt"

func main() {

	age := 21
	ageTwo := 20
	ageThree := 18


	if age <= 21 {
		fmt.Println("You're old enough to drink")
	}

	if ageTwo <= 20 {
		fmt.Println("You are not old enough to drink")
	}


	if ageThree > 12 && ageThree < 20 {
		fmt.Println("You are a teenager")
	}

}
