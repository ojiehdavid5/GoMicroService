package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
) 


func main(){

	//starting a server∏

	// fmt.Println("Hello, World!")
	// fmt.Println("This is a test.")
	port := 8080 
	 http.HandleFunc("/helloworld", helloWorldHandler) 
	  log.Printf("Server starting on port %v\n", port) 
	   http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	   	   log.Fatal()




	//Reading and writing json

	//“Marshalling Go structs to JSON”

	

	
	
}
//starting a server

type helloWorldResponse struct {     Message string  `json:"message"`  } 
type helloWorldRequest struct {     Name string `json:"name"`}

// type Requests struct { 
// 	 // Method specifies the HTTP method (GET, POST, PUT, etc.). 
// 	   Method string 
// // // Header contains the request header fields received by the server. The type Header is a link to map[string] []string. 
// //  Header Header 
// // // Body is the request's body.
//  Body io.ReadCloser 
//  } 



func helloWorldHandler(w http.ResponseWriter, r *http.Request) { 
	body,err:=io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w, "Bad request", http.StatusBadRequest)
return 
	}

	var request helloWorldRequest ;




	err = json.Unmarshal(body,&request)

	if err!=nil{
		http.Error(w, "Bad request", http.StatusBadRequest)
return 
	}

	
	// fmt.Fprint(w, "Hello World\n") 
	 response :=helloWorldResponse{Message:"Hello"    +  request.Name}
	 fmt.Println(response)

	encoder :=json.NewEncoder(w)
	encoder.Encode(response)


	//marshal 
	// data,err:=json.Marshal(response)
	// if err != nil {
	// 	panic("oops");

		


}
	// fmt.Fprint(w,string(data))
	// fmt.Print(string(data))




//reading and writing in go

// func Marshal(v interface{}) ([]byte, error) 

//unmarshal


