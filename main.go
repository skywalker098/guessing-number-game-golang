package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randGenerator() int {
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(100)
	return secret
}
func getValidInput(min, max int) int {
	var input int
	fmt.Println("Enter a number")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input")
		return getValidInput(min, max)
	} else if input < min || input > max {
		fmt.Println("Number out of range")
		return getValidInput(min, max)
	}
	return input
}

func main() {
	//getting input
	secret := randGenerator()
	fmt.Println(secret)
	// input := getValidInput(1, 100)

	attempt := 10

	for attempt > 0 {
		input := getValidInput(1, 100)
		if input == secret {
			fmt.Printf("You gussed the number in %d attempt and the number is %d", attempt, input)
			break
		} else if input > secret {
			fmt.Println("Input is high")
			attempt = attempt - 1
			continue
		} else if input < secret {
			fmt.Println("Input is low")
			attempt = attempt - 1

		}
	}

	//check whether the iniput is valid or not

	//compare with the secret number

	//response as per given input and decrease the chance
}
