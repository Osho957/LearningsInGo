package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1
	fmt.Println("Value of dice is: ",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Case 1: Random number is 1")
	case 2:
		fmt.Println("Case 2: Random number is 2")
		fallthrough
	case 3:
		fmt.Println("Case 3: Random number is 3")
	default:
		fmt.Println("Default case: Random number is unexpected")
	}
}