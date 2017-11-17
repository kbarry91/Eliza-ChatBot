/*
*	Author	: Kevin Barry
*	Date	: October 2017
*	Data Represenation and Querying Project
*	A chatbot web application using go based on the Eliza Program
*
*
 */

package main

import (
	"fmt"
	"net/http"

	"./eliza"
	//	"strings"
	//	"bytes"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// this is code that runs when a request is made to the /ask resource.
	userInput := r.URL.Query().Get("user-input")
	reply := eliza.AskEliza(userInput)
	fmt.Fprintf(w, reply)

} //chatHandler

func main() {

	//serve the files from the /static folder
	dir := http.Dir("./static")
	fileServer := http.FileServer(dir)
	/*code used for test purposes
	//newLines, _ := eliza.ReadLines("./database/responses.dat")
	//eliza.PrintResponses("./database/responses.dat")
	//fmt.Printf("%v", newLines)
	//tester := eliza.BuildResponses()
	//fmt.Print(tester)
	*/
	//handle requests to /
	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	http.ListenAndServe(":8080", nil)
}
