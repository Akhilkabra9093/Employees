package api

import (
	"Employees/internal"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func StartServer(addr string, router *gin.Engine) error {
	fmt.Printf("Server is running at %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		return err
	}
	return nil
}

func SetupRouter(db *sql.DB) *gin.Engine {
	// Initialize Gin router
	router := gin.Default()
	database := internal.GetDB()

	// Define API endpoints
	router.POST("/employees/paginated", func(c *gin.Context) {
		internal.ListEmployees(c, database)
	})

	router.POST("/employees", func(c *gin.Context) {
		// Parse JSON request body into Employee object
		var emp internal.Employee
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call CreateEmployee function with the parsed Employee object
		if err := internal.CreateEmployee(emp, database); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
	})

	router.GET("/employees/:id", func(c *gin.Context) {
		// Extract ID from URL parameters
		id := c.Param("id")
		employeeID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
		emp, err := internal.GetEmployee(employeeID, database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee"})
			return
		}

		if emp == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, emp)
	})

	router.PUT("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		employeeID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}

		var emp internal.Employee
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if err := internal.UpdateEmployee(employeeID, emp, database); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
	})

	router.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		employeeID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
		if err := internal.DeleteEmployee(employeeID, database); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	})

	return router
}
