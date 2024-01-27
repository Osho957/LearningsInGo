package main

import "fmt"


const LoginToken string = "jwt.123.xyz"

func main()  {
	var username string = "name"
	fmt.Println(username)
	fmt.Printf("Vairable is of type : %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Vairable is of type : %T \n",isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Vairable is of type : %T \n",smallVal)
	
	var smallFloat float32 = 255.4555858658765765
	fmt.Println(smallFloat)
	fmt.Printf("Vairable is of type : %T \n",smallFloat)
	
    // default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Vairable is of type : %T \n",anotherVariable)
	
	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("Vairable is of type : %T \n",anotherVariable1)
	
	//implicit type
	var website = "go.org"
	fmt.Println(website)

	// walrus operator (only allowed within a method)
	numberOfUser :=1000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Vairable is of type : %T \n",LoginToken)
	
}