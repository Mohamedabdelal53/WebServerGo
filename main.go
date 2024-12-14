package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	fmt.Println("Hello World")
	PortString := os.Getenv("PORT")
	if PortString == ""{
		log.Fatal("Port Not Found")
	}
	fmt.Println("Port:",PortString)
}
