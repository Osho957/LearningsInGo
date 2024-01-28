package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Intro to slices")

	var list = []string{"Apple", "Tomato"}
	fmt.Println(list)
	fmt.Printf("type of list is %T\n", list)

	list = append(list, "mango", "banana")
	fmt.Println(list)

	list = append(list[:3])
	fmt.Println(list)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 432
	highScores[2] = 980
	highScores[3] = 455
	// this will give error
	// highScores[4]= 986
	// but this will work slice property
	highScores = append(highScores, 435, 565, 123)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// sorting method
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	var courses = []string{"react","js","swift","python","java"}
	fmt.Println(courses)
    
	// removing from slices
    var idx int = 2;
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)
	
}
