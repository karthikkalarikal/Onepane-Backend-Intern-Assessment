package main

import (
	"fmt"
	"net/http"

	"github.com/karthikkalarikal/assignmentOnepane/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HandleCombinedResults)
	fmt.Println("Server is running on :8080")

	http.ListenAndServe(":8080", nil)
}
