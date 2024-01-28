package main

import "fmt"


func main()  {
	fmt.Println("structs in go")
	// no inheritence in golang; No super or parent
    user := User{"name","name@email.com",true,16}
	fmt.Println(user)
    fmt.Printf("user details : %+v\n", user)
    
	// accessing elements of class
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Println(user.Status)
	fmt.Println(user.Age)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int 
}