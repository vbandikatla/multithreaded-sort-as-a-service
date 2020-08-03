package api

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vbandikatla/multithreaded-sort-as-a-service/api/controllers"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("fetched env. values")
	}

	server.Initialize()
	server.Run()
}
