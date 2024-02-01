package main

import "fmt"

func main()  {
	fmt.Println("Loops in go lang")
	days := []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}
	fmt.Println(days)

	// way 1
	for d:=0; d<len(days); d++ {
		fmt.Println(days[d])
	}
     // way 2
	for i:= range days {
		fmt.Println(days[i])
	}

	// for each
	for index,day := range days{
		fmt.Printf("Index is %v and value is %v \n",index,day)
	}

	rougeValue :=1

	for rougeValue<10 {

		if rougeValue==2 {
			goto lco
		}
		
		if(rougeValue==5){
			rougeValue++
			continue
		}
		fmt.Println("Value is " , rougeValue)
		rougeValue++
	}

	lco:
	    fmt.Println("Go to lco")
		
}