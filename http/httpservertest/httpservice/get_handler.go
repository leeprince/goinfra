package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 10:13
 * @Desc:
 */

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(">>> getHandler")
		
		params := r.URL.Query()
		
		log.Println("Received params:", params)
		
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, world!",
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
