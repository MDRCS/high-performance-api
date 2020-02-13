package Routing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHttpGetAllEngineers(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpGetAllEngineers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"employee_id":2,"first_name":"John","last_name":"Jones","email":"jones.joe@gmail.com","gender":"M","employeer_id":{"employeer_id":2,"company_name":""}},{"employee_id":3,"first_name":"Ray","last_name":"Davis","email":"davis.ray@gmail.com","gender":"M","employeer_id":{"employeer_id":2,"company_name":""}}]`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestHttpGetOneEngineer(t *testing.T) {
	req, err := http.NewRequest("GET", "/engineer", bytes.NewBuffer([]byte("2")))
	if err != nil {
		t.Fatal(err)
	}

	req.URL,_ = url.Parse("http://localhost:8080/engineer/2")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpGetOneEngineer)
	handler.ServeHTTP(rr, req)
	t.Logf("url : %v ,  %v",req.URL,req.URL.Path)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"employee_id":0,"first_name":"","last_name":"","email":"","gender":"","employeer_id":{"employeer_id":0,"company_name":""}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHttpAddEngineer(t *testing.T) {

	var jsonStr = []byte(`{"first_name":"Anas","last_name":"Ben","email":"ben.anas@gmail.com","gender":"M","employeer_id":{"employeer_id":1,"company_name":""}}`)

	req, err := http.NewRequest("POST", "/engineer", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpAddEngineer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"first_name":"Anas","last_name":"Ben","email":"ben.anas@gmail.com","gender":"M","employeer_id":{"employeer_id":1,"company_name":""}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestHttpUpdateEngineer(t *testing.T) {

	var jsonStr = []byte(`{"first_name":"Anas","last_name":"Ben","email":"ben.anas@gmail.com","gender":"M","employeer_id":{"employeer_id":1,"company_name":""}`)

	req, err := http.NewRequest("PUT", "/engineer", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpUpdateEngineer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"first_name":"Anas","last_name":"Lhioui","email":"lhioui.anas@gmail.com","gender":"M","employeer_id":{"employeer_id":1,"company_name":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestHttpDeleteEngineer(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/engineer", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpDeleteEngineer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `the deletion has been complete successfully !!!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestHttpGetAllCompanies(t *testing.T) {
	req, err := http.NewRequest("GET", "/companies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpGetAllCompanies)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"employeer_id":1,"company_name":"Amazon"},{"employeer_id":2,"company_name":"Google"}]`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//func TestHttpGetOneCompany(t *testing.T) {
//
//}
//
//func TestHttpAddCompany(t *testing.T) {
//
//}
//
//func TestHttpDeleteCompany(t *testing.T) {
//
//}







