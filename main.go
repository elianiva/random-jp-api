package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

// TODO(elianiva): change this to `os.Getenv("PORT")` when we deploy this
var PORT = "3000"

func main() {
	mux := http.NewServeMux()

	n5 := make(map[string][]map[string]string)
	n5json, err := ioutil.ReadFile("./data/n5.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	json.Unmarshal(n5json, &n5)

	mux.HandleFunc("/api/n5/all", func(w http.ResponseWriter, r *http.Request) {
		jsonResp, err := json.Marshal(n5)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonResp))
	})

	mux.HandleFunc("/api/n5/rand", func(w http.ResponseWriter, r *http.Request) {
		randIdx := rand.Intn(len(n5["data"]))

		data := map[string]map[string]string{
			"data": n5["data"][randIdx],
		}

		jsonResp, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonResp))
	})

	fmt.Println("Server running on " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
