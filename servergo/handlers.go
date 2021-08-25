package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func hadleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CREATE")

}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "err :%v", err)
	}
	fmt.Fprintf(w, "PayLoad %v \n", metaData)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "err :%v", err)
		return
	}
	fmt.Println(user.mame)
	fmt.Fprintf(w, "PayLoad %v \n", user)
}
