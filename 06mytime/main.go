package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time Concepts ")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// using exact this format to get date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}