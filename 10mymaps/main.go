package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Maps in go lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("List Of All Languages: ", languages)
	fmt.Println("JS starts for: ", languages["JS"])
	
	//delete from map

	delete(languages,"RB")
	fmt.Println("List Of All Languages: ", languages)
	
	
	// for loops in golang
	
	for key,value :=range languages{
		fmt.Printf("for key %v, value is %v \n",key,value)
	}

}