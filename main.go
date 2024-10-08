package main

import (
	"arifptmx/helpers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := helpers.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// check if the connection work
	pingErr := db.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	// listen port
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
