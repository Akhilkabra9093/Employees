package internal

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListEmployees(c *gin.Context) {
	// Implement logic to list employees
	c.JSON(http.StatusOK, gin.H{
		"message": "List of employees",
	})
}

func CreateEmployee(emp Employee) error {
	_, err := Db.Exec("INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)", emp.Name, emp.Position, emp.Salary)
	if err != nil {
		log.Printf("Error inserting employee: %v", err)
	}
	return err
}

func GetEmployee(id int) (*Employee, error) {
	row := Db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = ?", id)
	emp := &Employee{}
	err := row.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("Error retrieving employee: %v", err)
	}
	return emp, err
}

func UpdateEmployee(id int, emp Employee) error {
	_, err := Db.Exec("UPDATE employees SET name=?, position=?, salary=? WHERE id=?", emp.Name, emp.Position, emp.Salary, id)
	if err != nil {
		log.Printf("Error updating employee: %v", err)
	}
	return err
}

func DeleteEmployee(id int) error {
	_, err := Db.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		log.Printf("Error deleting employee: %v", err)
	}
	return err
}
