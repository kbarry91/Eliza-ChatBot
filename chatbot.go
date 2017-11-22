/*
*	Author	: Kevin Barry
*	Date	: October 2017
*	Data Represenation and Querying Project
*	A chatbot web application using go based on the Eliza Program
 */

package main

import (
	"fmt"
	"net/http"
	// import eliza package
	"./eliza"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// this is code that runs when a request is made to the /Chat resource.
	userInput := r.URL.Query().Get("user-input")
	reply := eliza.AskEliza(userInput)
	fmt.Fprintf(w, reply)
} //chatHandler

func main() {

	//serve the files from the /static folder
	dir := http.Dir("./static")
	fileServer := http.FileServer(dir)
	//handle requests to /
	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)
	//listen on tcp and serve requests on port :8080
	http.ListenAndServe(":8080", nil)
}
