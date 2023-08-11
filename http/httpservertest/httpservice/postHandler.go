package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 10:12
 * @Desc:
 */

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println(">>> postHandler")
		
		var data struct {
			Data string `json:"data"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Received data " + data.Data,
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(response)
		return
	}
	
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
