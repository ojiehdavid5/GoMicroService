package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"log"

	"net/http"
)

// func path() {

// 	http.HandleFunc("/images/", newFooHandler)
// 	http.HandleFunc("/images/persian/", newBarHandler)
// 	http.HandleFunc("/images", newBuzzHandler)

// 	fmt.Println("Hello, World!")
// }
// func newFooHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello, World1!")

// }
// func newBarHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello, World!2")

// }
// func newBuzzHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello, World!3")

// }

func main() {
	// path()

	//starting a server∏

	// fmt.Println("Hello, World!")
	// fmt.Println("This is a test.")
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	// servering a static file to the broswer as response
	http.Handle("/images", http.FileServer(http.Dir("./images")))


	log.Printf("Server starting on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	log.Fatal()

	//Reading and writing json

	//“Marshalling Go structs to JSON”

}

//starting a server

type helloWorldResponse struct {
	Message string `json:"Message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// 	body,err:=io.ReadAll(r.Body)
	// 	if err!=nil{
	// 		http.Error(w, "Bad request", http.StatusBadRequest)
	// return
	// 	}

	var request helloWorldRequest
	fmt.Println(request)

	// err = json.Unmarshal(body,&request)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	fmt.Println(request)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return

	}

	// fmt.Fprint(w, "Hello World\n")
	response := helloWorldResponse{Message: "Hello" + request.Name}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// encoder.Encode(response)
	fmt.Println(response)

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
