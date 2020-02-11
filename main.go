package main

import (
	ht "./http_api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime"
)

func main() {

	//fmt.Println(db.Ping())
	//dbType := fmt.Sprintf("%T", db)
	//fmt.Println("Type of db object",dbType)

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	router := mux.NewRouter().StrictSlash(true)
	go router.HandleFunc("/", ht.HttpGetAllEngineers)
	go router.HandleFunc("/engineers", ht.HttpGetAllEngineers).Methods("GET")
	go router.HandleFunc("/engineers/{id}",ht.HttpGetOneEngineer).Methods("GET")
	go router.HandleFunc("/engineers",ht.HttpAddEngineer).Methods("POST")
	go router.HandleFunc("/engineers/{id}",ht.HttpUpdateEngineer).Methods("PATCH")
	go router.HandleFunc("/engineers/{id}",ht.HttpDeleteEngineer).Methods("DELETE")
	go router.HandleFunc("/companies",ht.HttpGetAllCompanies).Methods("GET")
	go router.HandleFunc("/companies/{name}",ht.HttpGetOneCompany).Methods("GET")
	go router.HandleFunc("/companies",ht.HttpAddCompany).Methods("POST")
	go router.HandleFunc("/companies/{id}",ht.HttpDeleteCompany).Methods("DELETE")

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
