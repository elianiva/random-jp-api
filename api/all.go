package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	level := map[string]JSONResult{
		"n5": File2GoMap(n5),
		"n4": File2GoMap(n4),
	}
	defer r.Body.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u := r.URL.Query()

	jsonResp, err := json.Marshal(level[u.Get("level")])
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResp))
	return
}
