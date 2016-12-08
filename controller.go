package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/init", initData).Methods("GET")
	router.HandleFunc("/test", getData).Methods("GET")
}

func initData(resp http.ResponseWriter, req *http.Request) {

	insert()
}

func getData(resp http.ResponseWriter, req *http.Request) {

	val, _ := json.Marshal(callDb())

	resp.Write(val)
}
