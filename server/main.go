package main

import (
	//"fmt"
	"io"
	"log"
	"net/http"
	"orderApp/server/routes"
	"os"
)

func main() {
	logFileName := os.Args[0] + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Panicln(err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			panic(r)
		}
	}()

	routes.Routes()

	log.Println("Starting server...")
	err = http.ListenAndServe(":9090", nil) // set listen port
	log.Println("...Started")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
