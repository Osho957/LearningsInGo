package main

import "fmt"

func main() {
	// it executes at last
	defer fmt.Println("Hello")
	defer fmt.Println("Hello World")
	defer fmt.Println("Hello 2")
	defer fmt.Println()
	fmt.Println("welcome to defer")
	myDefer()

}

func myDefer(){
	for i:=0; i<5; i++ {
		defer fmt.Print(i)
	} 
	
}