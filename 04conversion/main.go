package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	fmt.Println("Provide Rating to app")

	reader :=bufio.NewReader(os.Stdin)

	input, _ :=reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err!=nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("Changed to float: ",numRating+1)
	}
	
}