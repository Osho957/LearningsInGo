package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array in go")

	// put size and datatype
	var list[4] string;

	list[0] = "hello"
	list[1] = "value"
	list[3] = "bny"

	fmt.Println("list is : ",list)
	fmt.Println("length is : ",len(list))

	var veg = [5]string{"potato", "tomato", "mushroom"}
	fmt.Println("veg list is: ",veg)
}

