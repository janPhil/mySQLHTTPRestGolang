package models

import (
	"database/sql"
	"log"
)

// Employee is the struct used for response
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Salary    int
}

// AllEmployees returns all employees found in the database
func AllEmployees(db *sql.DB) ([]*Employee, error) {

	employees := make([]*Employee, 0)
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		tmp := new(Employee)
		err := rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Salary)
		employees = append(employees, tmp)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
