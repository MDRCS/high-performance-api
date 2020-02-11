package http_api

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HttpGetAllEngineers(writer http.ResponseWriter, request *http.Request) {

	start := time.Now()
	employees := GetEngineers()
	buffer, err := json.Marshal(employees)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s ",string(buffer))
	log.Println(time.Since(start))
}

func HttpGetOneEngineer(writer http.ResponseWriter,request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	empID,_ := strconv.Atoi(id)
	employees := GetOneEngineer(empID)
	buffer, err := json.Marshal(employees)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s ",string(buffer))
	log.Println(time.Since(start))
}

func HttpAddEngineer(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	new_employee := Employee{}
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_employee)
	Insert(new_employee)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_employee)
	log.Println(time.Since(start))

}

func HttpUpdateEngineer(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	new_employee := Employee{}
	id := mux.Vars(request)["id"]
	empID,_ := strconv.Atoi(id)
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_employee)
	Update(empID,&new_employee)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_employee)
	log.Println(time.Since(start))

}

func HttpDeleteEngineer(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	empID, _ := strconv.Atoi(id)
	Delete(empID)
	fmt.Fprintf(writer, "the deletion has been complete successfully !!! ")
	log.Println(time.Since(start))
}

func HttpGetAllCompanies(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	companies := GetCompanies()
	buffer, err := json.Marshal(companies)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s ",string(buffer))
	log.Println(time.Since(start))

}

func HttpGetOneCompany(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	camp_name := mux.Vars(request)["name"]
	campany := GetOneCompanyByName(camp_name)
	buffer, err := json.Marshal(campany)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s",string(buffer))
	log.Println(time.Since(start))

}

func HttpAddCompany(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	new_campany := Employeer{}
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_campany)
	AddCompany(new_campany)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_campany)
	log.Println(time.Since(start))
}


func HttpDeleteCompany(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	campID, _ := strconv.Atoi(id)
	DeleteCompany(campID)
	fmt.Fprintf(writer, "the deletion has been complete successfully !!! ")
	log.Println(time.Since(start))
}

func HttpCheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ValidNewRecord(err error,writer http.ResponseWriter) {
	if err != nil {
		fmt.Fprintf(writer,"Oups there is a problem with new record's data ..")
	}
}