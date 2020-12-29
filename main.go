package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var PORT = os.Getenv("PORT")

type JSONResult map[string][]map[string]string

func File2GoMap(path string) JSONResult {
	result := make(JSONResult)
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	json.Unmarshal(jsonFile, &result)
	return result
}

func main() {
	level := map[string]JSONResult{
		"n5": File2GoMap("./data/n5.json"),
		"n4": File2GoMap("./data/n5.json"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Hello 世界！")
	})

	http.HandleFunc("/api/all", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		u := r.URL.Query()

		jsonResp, err := json.Marshal(level[u.Get("level")])
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonResp))
	})

	http.HandleFunc("/api/rand", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Access-Control-Allow-Origin", "*")

		u := r.URL.Query()
		randIdx := rand.Intn(len(level[u.Get("level")]["data"]))

		data := map[string]map[string]string{
			"data": level[u.Get("level")]["data"][randIdx],
		}

		jsonResp, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonResp))
	})

	log.Println("Server started on port: " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
