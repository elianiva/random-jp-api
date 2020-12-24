package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO(elianiva): change this to `os.Getenv("PORT")` when we deploy this
var PORT = "3000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello 世界")
	})

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
	fmt.Println("Server running on " + PORT)
}
