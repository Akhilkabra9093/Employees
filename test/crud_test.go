package test

import (
	"Employees/internal"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emp := internal.Employee{
		Name:     "Akhil test",
		Position: "Intern",
		Salary:   600000,
	}

	mock.ExpectExec("INSERT INTO employees").
		WithArgs(emp.Name, emp.Position, emp.Salary).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = internal.CreateEmployee(emp, db)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emp := &internal.Employee{
		ID:       1,
		Name:     "Akhil test 2",
		Position: "Associate",
		Salary:   6000,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "position", "salary"}).
		AddRow(emp.ID, emp.Name, emp.Position, emp.Salary)

	mock.ExpectQuery("SELECT id, name, position, salary FROM employees WHERE id = ?").
		WithArgs(emp.ID).
		WillReturnRows(rows)

	result, err := internal.GetEmployee(emp.ID, db)
	assert.NoError(t, err)
	assert.Equal(t, emp, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEmployee(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create an instance of the employee to be updated
	emp := internal.Employee{
		ID:       1,
		Name:     "Akhil New Test",
		Position: "Senior Developer",
		Salary:   8000,
	}

	// Define the expected SQL query and its arguments
	mock.ExpectExec("UPDATE employees SET name=?, position=?, salary=? WHERE id=?").
		WithArgs(emp.Name, emp.Position, emp.Salary, emp.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the UpdateEmployee function with the mock database connection
	err = internal.UpdateEmployee(emp.ID, emp, db)
	if err != nil {
		t.Errorf("UpdateEmployee returned an unexpected error: %v", err)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	empID := 1

	mock.ExpectExec("DELETE FROM employees WHERE id=?").
		WithArgs(empID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = internal.DeleteEmployee(empID, db)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
