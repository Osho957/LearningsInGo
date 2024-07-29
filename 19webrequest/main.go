package main

import (
	"fmt"
	"io"
	"net/http"
) 

const url = "https://lco.dev"

func main() {
	fmt.Println("handle request")
    response,err := http.Get(url)
	if(err!=nil){
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n",response)
	fmt.Println(response.StatusCode)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if(err!=nil){
		panic(err)
	}
	content  := string(databytes)
	fmt.Println(content)
}