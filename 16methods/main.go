package main

import "fmt"


func main()  {
	fmt.Println("methods in go")
	// no inheritence in golang; No super or parent
    user := User{"name","name@email.com",true,16}
	fmt.Println(user)
    fmt.Printf("user details : %+v\n", user)
    
	// accessing elements of class
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	user.GetStatus()
	fmt.Println(user.Age)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int 
}

func (u User) GetStatus()  {
	fmt.Println("is user active: ", u.Status)
}