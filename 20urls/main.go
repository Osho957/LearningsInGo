package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("urls handling")
    fmt.Println(myurl)

	//parsing url
	result,err := url.Parse(myurl)
	if(err!=nil){
		panic(err)
	}
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)
	fmt.Println(qparams)

	for _,val :=range qparams{
		fmt.Println("Param is : ",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=osho",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}