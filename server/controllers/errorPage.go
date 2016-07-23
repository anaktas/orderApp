package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside of ErrorNotFound")
	fmt.Fprintf(w, "404: Not Found") // send data to client side
}
