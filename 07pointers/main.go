package main

import "fmt"

func main()  {
	fmt.Println("welcome to a class on pointer")

	// default value is nill
	var ptr *int 
	fmt.Println("Value of pointer is ",ptr)

	myNumber := 23

	// refernce means &
	var ptr2 = &myNumber

	// it prints actual memory address
	fmt.Println("Value of ptr is ",ptr2)

	// it will print the exact value
	fmt.Println("Value of ptr value is ", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("Value of ptr value is ", *ptr2)

}