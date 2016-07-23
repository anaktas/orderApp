package waiter

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside of waiter.Index")
	fmt.Fprintf(w, "Hello Taso! You are in the waiter index") // send data to client side
}
