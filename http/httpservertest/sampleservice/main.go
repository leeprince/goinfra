package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port *int
)

func main() {
	port = flag.Int("port", 8090, "port")
	flag.Parse()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/prince/get", getHandler)
	http.HandleFunc("/prince/post", postHandler)
	fmt.Printf("Server listening on port:%d ...\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<head>
			    <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
			    <script>
			        function sendGet() {
			            $.get("/prince/get", function (data, status) {
			                console.log("/prince/get", data, status)
			            });
			        }
			
			        function sendPost() {
			            $.ajax({
			                url: "/prince/post",
			                type: 'POST',
			                aysnc: true,
			                dataType: "json",
			                contentType: 'application/json;charset=UTF-8',
			                data: JSON.stringify({
			                    name: "Donald Duck",
			                    city: "Duckburg"
			                }),
			                success: function (r) {
			                    console.log("/prince/post", r)
			                }
			            })
			        }
			    </script>
			</head>
			<body>
			<form>
			    <label for="input">Input:</label>
			    <input id="input" name="input" value="">
			    <p>
			
			        <label for="checkbox1">Checkbox 1:</label>
			        <input type="checkbox" id="checkbox1" name="checkbox1" value="1">
			        <label for="checkbox2">Checkbox 2:</label>
			        <input type="checkbox" id="checkbox2" name="checkbox2" value="2">
			    <p>
			
			        <label for="select1">Select:</label>
			        <select id="select1" name="select1">
			            <option value="1">Option 1</option>
			            <option value="2">Option 2</option>
			            <option value="3">Option 3</option>
			        </select>
			    <p>
			
			        <label for="select2">Select multiple:</label>
			        <select id="select2" name="select2" multiple>
			            <option value="1">Option 1</option>
			            <option value="2">Option 2</option>
			            <option value="3">Option 3</option>
			        </select>
			    <p>
			        <label for="textarea">Textarea:</label>
			        <textarea id="textarea" name="textarea" style="width:200px;height:50px">textarea</textarea>
			    <p>
			
			        <input type="submit" value="Submit">
			</form>
			
			<hr>
			
			<input id="sendGet" type="button" value="Send GET" onclick="sendGet()">
			<input id="sendPost" type="button" value="Send POST" onclick="sendPost()">
			</body>
			</html>
		`)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

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
