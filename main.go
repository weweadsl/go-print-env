package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()

	version := os.Getenv("Version")
	secretKey := os.Getenv("SECRET_KEY")
	cloudP := os.Getenv("CLOUDSDK_PYTHON")

	fmt.Println("Bucket: ",version)
	fmt.Println("secretKey: ",secretKey)
	fmt.Println("cloudPython: ", cloudP)
}