package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//

type CourseDB struct {
	db        *sql.DB
	COURSE_ID int
}

var database *CourseDB

func OpenDB() (err error) {
	database = new(CourseDB)

	database.db, err = sql.Open("sqlite3", "./w3e01.db")
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//

func (db *CourseDB) createTables() {
	s := "CREATE TABLE course ( id INT NOT NULL PRIMARY KEY, name VARCHAR(255) NOT NULL, description VARCHAR(255))"
	if _, err := database.db.Exec(s); err != nil {
		if err.Error() == "table course already exists" {
			database.db.Exec("DROP TABLE course")
			database.createTables()
			return
		}
		log.Println("createTables():", err.Error())
	}

}

func (db *CourseDB) addCourse(name, description string) {
	var err error
	var tx *sql.Tx
	var stmt *sql.Stmt

	db.COURSE_ID++
	query := "INSERT INTO course VALUES ( ? , ?, ?)"

	tx, _ = database.db.Begin()
	if stmt, err = tx.Prepare(query); err != nil {
		log.Println("addCourse() stmt:", err.Error())
		tx.Rollback()
		return
	}
	defer stmt.Close()

	if _, err = stmt.Exec(db.COURSE_ID, name, description); err != nil {
		log.Println("addCourse() exec:", err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
}

func (db *CourseDB) listCourses() {
	var course *Course

	query := "SELECT * FROM course"
	rows, _ := database.db.Query(query)
	defer rows.Close()

	for rows.Next() {
		course = new(Course)
		rows.Scan(&course.id, &course.name, &course.description)
		fmt.Println(course)
	}
	rows.Close()

}

func (db *CourseDB) listCoursesByName(name string) {
	var err error
	var stmt *sql.Stmt
	var rows *sql.Rows
	var course *Course

	query := "SELECT * FROM course WHERE name = ?"

	stmt, _ = database.db.Prepare(query)
	if rows, err = stmt.Query(name); err != nil {
		log.Println("listCoursesByName():", err.Error())
		return
	}
	for rows.Next() {
		course = new(Course)
		rows.Scan(&course.id, &course.name, &course.description)
		fmt.Println(course)
	}
	rows.Close()
}

//

type Course struct {
	id          int
	name        string
	description string
}

func (c *Course) String() string {
	return fmt.Sprintf("%d\t%s\t%s", c.id, c.name, c.description)
}

//

func main() {
	OpenDB()
	database.createTables()

	fmt.Println("COURSES:")
	database.listCourses()

	database.addCourse("ohpe", "ohjelmoinnin perusteet")

	fmt.Println("---")
	database.listCourses()

	database.addCourse("ohja", "ohjelmoinnin jatkokurssi")
	fmt.Println("---")
	database.listCourses()

	database.addCourse("wepa", "web-palvelinohjelmointi")

	fmt.Println("---")
	database.listCourses()

	fmt.Println("---")
	database.listCoursesByName("wepa")
}
