package Database

import (
	"database/sql"
	c "../Credentials"
	m "../Models"
)

func GetConn(user, password, dbname  string) sql.DB{
	db, err := sql.Open("mysql", user+ ":"+ password+"@/"+dbname)
	CheckError(err)
	return *db
}


func GetEngineers() []m.Employee {

	var employees []m.Employee
	var emp m.Employee
	var employeer m.Employeer

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	result,err := db.Query("select * from employee")
	CheckError(err)

	for result.Next() {

		err := result.Scan(&emp.ID,&emp.Firstname,&emp.Lastname,&emp.Email,&emp.Gender,&employeer.ID)
		emp.Employeer = employeer
		employees = append(employees,emp)
		CheckError(err)
	}

	defer db.Close()

	return employees
}

func GetOneEngineer(empID int) m.Employee {

	var emp m.Employee
	var employeer m.Employeer

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	result,err := db.Query("select * from employee where employee_id = ?",empID)
	CheckError(err)

	for result.Next() {
		err := result.Scan(&emp.ID,&emp.Firstname,&emp.Lastname,&emp.Email,&emp.Gender,&employeer.ID)
		emp.Employeer = employeer
		CheckError(err)
	}

	defer db.Close()

	return emp
}

func Update(emp_id int, emp *m.Employee) {

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	_, err := db.Query("update employee set email=? ,employeer_id=? where employee_id=?",emp.Email,emp.Employeer.ID,emp_id)
	CheckError(err)
	defer db.Close()
}

func Delete(emp_id int){

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	_, err := db.Query("delete from employee where employee_id =?",emp_id)
	CheckError(err)
	defer db.Close()
}

func DeleteByCamp(camp_id int){

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	_, err := db.Query("delete from employee where employeer_id =?",camp_id)
	CheckError(err)
	defer db.Close()
}


func Insert(emp m.Employee) {

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	_, err := db.Query("insert into employee " +
			"(first_name,last_name,email,gender,employeer_id) values (?,?,?,?,?)",
			emp.Firstname,
			emp.Lastname,
			emp.Email,
			emp.Gender,
			emp.Employeer.ID)

	CheckError(err)
	defer db.Close()

}


func AddCompany(camp m.Employeer) {
	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	_, err := db.Query("insert into employeer " +
		"(company_name) values (?)",
		camp.Company)
	CheckError(err)
	defer db.Close()
}

func GetCompanies() []m.Employeer {

	employeer := m.Employeer{}
	companies := []m.Employeer{}

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	result,err := db.Query("select * from employeer")
	CheckError(err)
	for result.Next() {
		err := result.Scan(&employeer.ID,&employeer.Company)
		companies = append(companies,employeer)
		CheckError(err)
	}

	defer db.Close()
	return companies
}

func GetOneCompanyByName(name string) m.Employeer {

	camp := m.Employeer{}
	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	result,err := db.Query("select * from employeer where company_name  = ?",name)
	CheckError(err)

	for result.Next() {
		err:= result.Scan(&camp.ID,&camp.Company)
		CheckError(err)
	}

	defer db.Close()

	return camp
}

func DeleteCompany(camp_id int){

	db := sql.DB{}
	db = GetConn(c.Mysql_user,c.Mysql_password,c.Database_name)
	DeleteByCamp(camp_id) //Cascade erasing
	_, err := db.Query("delete from employeer where employeer_id =?",camp_id)
	CheckError(err)
	defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}