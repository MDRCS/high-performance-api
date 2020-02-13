package app

import (
	rt "../../http_api/Routing"
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
	go router.HandleFunc("/", rt.HttpGetAllEngineers)
	go router.HandleFunc("/engineers", rt.HttpGetAllEngineers).Methods("GET")
	go router.HandleFunc("/engineer/{id}",rt.HttpGetOneEngineer).Methods("GET")
	go router.HandleFunc("/engineer",rt.HttpAddEngineer).Methods("POST")
	go router.HandleFunc("/engineer/{id}",rt.HttpUpdateEngineer).Methods("PATCH")
	go router.HandleFunc("/engineer/{id}",rt.HttpDeleteEngineer).Methods("DELETE")
	go router.HandleFunc("/companies",rt.HttpGetAllCompanies).Methods("GET")
	go router.HandleFunc("/company/{name}",rt.HttpGetOneCompany).Methods("GET")
	go router.HandleFunc("/company",rt.HttpAddCompany).Methods("POST")
	go router.HandleFunc("/company/{id}",rt.HttpDeleteCompany).Methods("DELETE")

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
