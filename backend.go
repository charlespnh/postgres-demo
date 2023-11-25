package main

import (
	"database/sql"
	"fmt"
	// "os"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB // created outside to make it global.

var (
	HOST = "localhost"
	PORT = 5432
	USER = "postgres"
	DBNAME = "postgres"
	PASSWORD = "postgres"
)

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {
	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		HOST, PORT, USER, DBNAME, PASSWORD)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}

/**************************************************************************************/

func AddStudent(student Student) (int64, error) {
	// insert into students (first_name, last_name, email, enrollment_date)
	// values ('John', 'Doe', 'johndoe@gmail', '2021-01-01');
	sqlStatement := `
	INSERT INTO students (first_name, last_name, email, enrollment_date)
	VALUES ($1, $2, $3, $4);`
	res, err := Db.Exec(sqlStatement, student.first_name, student.last_name, student.email, student.enrollment_date)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	// fmt.Println(count)
	return count, nil
}

func GetAllStudents() ([]Student, error) {
	students := []Student{}
	sqlStatement := `SELECT * FROM students;`
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		student := Student{}
		err = rows.Scan(&student.student_id, &student.first_name, &student.last_name, &student.email, &student.enrollment_date)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
		// fmt.Println(student)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func UpdateStudentEmail(student_id int, email string) (int64, error) {
	sqlStatement := `
	UPDATE students
	SET email = $2
	WHERE student_id = $1;`
	res, err := Db.Exec(sqlStatement, student_id, email)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	// fmt.Println(count)
	return count, nil
}

func DeleteStudent(student_id int) (int64, error) {	
	sqlStatement := `
	DELETE FROM students
	WHERE student_id = $1;`
	res, err := Db.Exec(sqlStatement, student_id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	// fmt.Println(count)
	return count, nil
}

/**************************************************************************************/

/*
func CreateDemoTable() (int64, error) {
	
	sqlSchema := `
	CREATE TABLE IF NOT EXISTS students (
		student_id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		enrollment_date DATE
	);`
	_, err := Db.Exec(sqlSchema)
	if err != nil {
		return 0, err
	}

	// insert into students (first_name, last_name, email, enrollment_date)
	// values ('John', 'Doe', 'johndoe@gmail', '2021-01-01');
	students = []Student{
		{first_name: "John", last_name: "Doe", email: "johndoe@gmail", enrollment_date: "2023-09-01"},
		{first_name: "Jane", last_name: "Smith", email: "janesmith@gmail", enrollment_date: "2023-09-01"},
		{first_name: "Jim", last_name: "Bim", email: "bobsmith@gmail", enrollment_date: "2021-09-02"},
	}

	for _, student := range students {
		sqlStatement := `
		INSERT INTO students (first_name, last_name, email, enrollment_date)
		VALUES ($1, $2, $3, $4);`
		_, err := Db.Exec(sqlStatement, student.first_name, student.last_name, student.email, student.enrollment_date)
		if err != nil {
			return 0, err
		}
		fmt.Println("New record ID is:", id)
	}
}
*/
