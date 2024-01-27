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

	createdDate := time.Date(2001,time.May,06,06,06,00,0,time.UTC)
	fmt.Println("Created At: ",createdDate)
	fmt.Println(createdDate.Format("01-02-2006"))

	var local = time.Now().Nanosecond()
	fmt.Println(local)
}

/*
go build -> to build executable files
GOOS = "windows" for window os
       "linux" for linux
*/