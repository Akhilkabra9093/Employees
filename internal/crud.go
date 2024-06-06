package internal

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListEmployees(c *gin.Context, db *sql.DB) {
	var req PaginationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}

	offset := (req.Page - 1) * req.Size
	rows, err := db.Query("SELECT id, name, position, salary FROM employees LIMIT ? OFFSET ?", req.Size, offset)
	if err != nil {
		log.Printf("Error fetching employees: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
			log.Printf("Error scanning employee: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
			return
		}
		employees = append(employees, emp)
	}

	// Get the total number of employees
	var totalCount int
	err = db.QueryRow("SELECT COUNT(*) FROM employees").Scan(&totalCount)
	if err != nil {
		log.Printf("Error counting employees: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees count"})
		return
	}

	// Respond with paginated employees and metadata
	c.JSON(http.StatusOK, gin.H{
		"page":      req.Page,
		"size":      req.Size,
		"total":     totalCount,
		"employees": employees,
	})
}

func CreateEmployee(emp Employee, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)", emp.Name, emp.Position, emp.Salary)
	if err != nil {
		log.Printf("Error inserting employee: %v", err)
	}
	return err
}

func GetEmployee(id int, db *sql.DB) (*Employee, error) {
	row := db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = ?", id)
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

func UpdateEmployee(id int, emp Employee, db *sql.DB) error {
	_, err := db.Exec("UPDATE employees SET name=?, position=?, salary=? WHERE id=?", emp.Name, emp.Position, emp.Salary, id)
	if err != nil {
		log.Printf("Error updating employee: %v", err)
	}
	return err
}

func DeleteEmployee(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		log.Printf("Error deleting employee: %v", err)
	}
	return err
}
