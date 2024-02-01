package main

import "fmt"

func main()  {
	fmt.Println("functions")
	greeting();
    result := add(2,3)
	fmt.Println(result)

	total := proAdd(1,2,2,3,4,5,5,5,5)
	fmt.Println(total)
}

func add(x int, y int) int{
	return x+y;
}
func proAdd(values ...int) int{
	total :=0
    for _,val := range(values){
		total+=val
	}
	return total
}
func greeting(){
	fmt.Println("Greet")
}