package main

import (
	"fmt"
)

func exercise_5() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}
}

func exercise_6() {
	x := 5
	if x < 50 {
		fmt.Println("x is less than 50")
	} else {
		fmt.Println("x is greater than 50")
	}
}

func exercise_7() {
	x := 10
	if x < 100 {
		fmt.Println("x is less than 100")
	} else if x == 10 {
		fmt.Println("x is greater than 100")
	}
}

func exercise_8() {
	x := 10

	switch {
	case x == 10:
		fmt.Println("x is equal to 10")
	case x < 100:
		fmt.Println("x is less than 100")
	default:
		fmt.Println("x is greater than 100")

	}
}

func exercise_9() {
	favoriteSport := "Boxe"
	switch favoriteSport {
	case "Boxe":
		fmt.Println("massa! boxe é uma luta interessante")
	case "Futebol":
		fmt.Println("classico!")
	}
}
