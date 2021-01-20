package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONResult map[string][]map[string]string

func File2GoMap( /*path*/ data string) JSONResult {
	result := make(JSONResult)
	// jsonFile, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil
	// }

	json.Unmarshal([]byte(data), &result)
	return result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello 世界！")
	return
}
