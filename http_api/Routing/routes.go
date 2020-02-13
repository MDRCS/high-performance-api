package Routing

import (
	db "../Database"
	m "../Models"
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
	employees := db.GetEngineers()
	buffer, err := json.Marshal(employees)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s",string(buffer))
	log.Println(time.Since(start))
}

func HttpGetOneEngineer(writer http.ResponseWriter,request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	empID,_ := strconv.Atoi(id)
	employees := db.GetOneEngineer(empID)
	buffer, err := json.Marshal(employees)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s",string(buffer))
	log.Println(time.Since(start))
}

func HttpAddEngineer(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	new_employee := m.Employee{}
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_employee)
	db.Insert(new_employee)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_employee)
	log.Println(time.Since(start))

}

func HttpUpdateEngineer(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	new_employee := m.Employee{}
	id := mux.Vars(request)["id"]
	empID,_ := strconv.Atoi(id)
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_employee)
	db.Update(empID,&new_employee)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_employee)
	log.Println(time.Since(start))

}

func HttpDeleteEngineer(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	empID, _ := strconv.Atoi(id)
	db.Delete(empID)
	fmt.Fprintf(writer, "the deletion has been complete successfully !!!")
	log.Println(time.Since(start))
}

func HttpGetAllCompanies(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	companies := db.GetCompanies()
	buffer, err := json.Marshal(companies)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s",string(buffer))
	log.Println(time.Since(start))

}

func HttpGetOneCompany(writer http.ResponseWriter,request *http.Request) {

	start := time.Now()
	camp_name := mux.Vars(request)["name"]
	campany := db.GetOneCompanyByName(camp_name)
	buffer, err := json.Marshal(campany)
	HttpCheckError(err)
	fmt.Fprintf(writer,"%s",string(buffer))
	log.Println(time.Since(start))

}

func HttpAddCompany(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	new_campany := m.Employeer{}
	record,err := ioutil.ReadAll(request.Body)
	ValidNewRecord(err,writer)
	json.Unmarshal(record,&new_campany)
	db.AddCompany(new_campany)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(new_campany)
	log.Println(time.Since(start))
}


func HttpDeleteCompany(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	id := mux.Vars(request)["id"]
	campID, _ := strconv.Atoi(id)
	db.DeleteCompany(campID)
	fmt.Fprintf(writer, "the deletion has been complete successfully !!!")
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