package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func User(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	
	id := r.URL.Query().Get("id")
	resp["id"] = id
	currentPath := strings.TrimPrefix(r.URL.Path, "/user/")
	resp["path"] = currentPath
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}