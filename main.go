package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
