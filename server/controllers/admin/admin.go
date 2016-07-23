package admin

import (
	"fmt"
	"log"
	"net/http"
	//"orderApp/db"
)

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside of admin.Index")
	//db.ConnectionToMySQL()
	fmt.Fprintf(w, "Hello Taso! You are in the admin index") // send data to client side
}
