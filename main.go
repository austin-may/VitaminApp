package main

import (
	"VitaminApp/database"
	"VitaminApp/vitamin"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//will probably deprecate soon
	database.SetupDatabase()
	vitamin.SetupRoutes()
	fmt.Println("Listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
