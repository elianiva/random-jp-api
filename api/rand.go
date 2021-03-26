package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Rand(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	level := map[string]JSONResult{
		"n5": File2GoMap(n5),
		"n4": File2GoMap(n4),
	}
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
	return
}
