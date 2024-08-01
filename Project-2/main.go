package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		randomNumber := rand.Intn(100) + 1
		var estimate int

		fmt.Println("Keep in your mind a number which is between 1- 100. Can u find it?")

		for {
			fmt.Print("What is your estimate? ")
			fmt.Scan(&estimate)

			if estimate < randomNumber {
				fmt.Println("It is a bigger number!")
			} else if estimate > randomNumber {
				fmt.Println("It is a lower number!")
			} else {
				fmt.Println("Yeah Perfect! This is True.")
				break
			}
		}

		var choice int
		fmt.Println("Do u wanna play it again ? (1: Yeah, 2: Nope)")
		fmt.Scan(&choice)

		if choice != 1 {
			fmt.Println("Game Over. Thanks for Playing!")
			break
		}
	}
}
