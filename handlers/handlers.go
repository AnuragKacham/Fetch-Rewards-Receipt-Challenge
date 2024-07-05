package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var dataStore = make(map[string]string)

type CustomData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetCustomData(w http.ResponseWriter, r *http.Request) {
	var data CustomData
	_ = json.NewDecoder(r.Body).Decode(&data)
	dataStore[data.Key] = data.Value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func GetCustomData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, exists := dataStore[key]
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}
