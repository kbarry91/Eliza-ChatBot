
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
	"net/http"
	"fmt"
//	"strings"
//	"bytes"
)
	


func chatHandler(w http.ResponseWriter, r *http.Request) {
	

}//chatHandler

func main() {
	//serve the files from the /static folder
	dir := http.Dir("./static")
	fileServer := http.FileServer(dir)
	//handle requests to /
	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	http.ListenAndServe(":8080", nil)
}
