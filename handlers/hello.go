package handlers

import (
	"fmt"
	"net/http"
)

//Hello serves route to print hello to std out
func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello")
}