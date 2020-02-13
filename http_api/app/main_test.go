package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"os"
	"testing"
	c "../Credentials"
)

const tableCreationQueryEmployee = `
CREATE TABLE IF NOT EXISTS employee
(
    employee_id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(20) NOT NULL,
	last_name VARCHAR(20) NOT NULL,
	email VARCHAR(30) NOT NULL,
	gender VARCHAR(1) NOT NULL,
    PRIMARY KEY (OrderID),
    FOREIGN KEY (employeer_id) REFERENCES employeer(employeer_id)
)`


const tableCreationQueryEmployeer =`
CREATE TABLE IF NOT EXISTS employeer
(
	employeer_id INT AUTO_INCREMENT NOT NULL,
	company_name VARCHAR(20),
)`


func TestMain(m *testing.M) {
	mux.NewRouter()
	ensureTableExists()
	main()
	clearTable()
	os.Exit(1)
}

func ensureTableExists() {

	DB,_ := sql.Open("mysql", c.Mysql_user+ ":"+ c.Mysql_password+"@/"+c.Database_name)

	if _, err := DB.Exec(tableCreationQueryEmployee); err != nil {
		log.Fatal(err)
	}

	if _, err := DB.Exec(tableCreationQueryEmployeer); err != nil {
		log.Fatal(err)
	}
}


func clearTable() {
	DB,_ := sql.Open("mysql", c.Mysql_user+ ":"+ c.Mysql_password+"@/"+c.Database_name)
	DB.Exec("DELETE FROM employee")
	DB.Exec("ALTER TABLE employee AUTO_INCREMENT = 1")
}



