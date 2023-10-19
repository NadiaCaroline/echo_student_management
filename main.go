package main

import (
	"net/http" // interaksi dengan hhtp
	"strconv"  // string convertion

	"github.com/labstack/echo/v4" // inisiasi framework echo
)

// inisiasi Constructur
type Student struct { // struct dari data yang ada (data yang mau dipakai)

		ID    int    `json:"id"` // specify untuk hasil datanya berbentuk json dan keynya ID
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Grade string `json:"grade"`
	}
	
// Inisiasi Database 
var students []Student

func main() {
	// inisiasi echo framework 
	e := echo.New()

	//Routing
	e.GET("/students", getStudents) // ambil seluruh data
	e.GET("/students/:id", getStudent) // ambil sesuai ID
	e.POST("/students", createStudent) // bikin student
	e.PUT("/students/:id", updateStudent) // update student
	e.DELETE("/students/:id", deleteStudent) // delete student

	// inisiasi port
	e.Start(":8080")
}

// Mengambil semua data yang ada dan mengembalikan HTTP STATUS: GET
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
} 
// Mengambil data student berdasarkan ID dan mengembalikan Status: GET
func getStudent(c echo.Context) error {
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id")) // method strconv untuk mengambil information dari ID dan casting menjadi integer
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

// Membuat data baru: POST
func createStudent(c echo.Context) error {
	// inisiasi Student baru
	student := new(Student)

	// Mengembalikan informasi student
	if err := c.Bind(student); err != nil {
		return err
	}

	// ID Baru dan menambahkan data kedalam database
	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

// update data student: PUT
func updateStudent(c echo.Context) error{
	// mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			updatedStudent := new(Student)
			if err := c.Bind(updatedStudent ); err != nil {
				return err
			}

			// update student data
			students[i].Name = updatedStudent.Name
			students[i].Age  = updatedStudent.Age
			students[i].Grade = updatedStudent.Grade

			return c.JSON(http.StatusOK, students[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

// Menghapus informasi data student: DELETE

func deleteStudent(c echo.Context) error{
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:] ... )
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

