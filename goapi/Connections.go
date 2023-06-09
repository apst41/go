package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "Paytm@197",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "ajay",
		AllowNativePasswords: true,
	}

	// Replace the connection parameters with your MySQL configuration

	// Open a connection to the database
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetAjay(db *sql.DB, id int64) (User, error) {
	row := db.QueryRow("SELECT id, name, age FROM ajay WHERE id = ?", id)
	var ajay User
	err := row.Scan(&ajay.ID, &ajay.Name, &ajay.Age)
	if err != nil {
		return User{}, err
	}
	return ajay, nil
}

func FetchAllRows(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM ajay")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []User

	for rows.Next() {
		var ajay User
		err := rows.Scan(&ajay.ID, &ajay.Name, &ajay.Age)
		if err != nil {
			return nil, err
		}
		results = append(results, ajay)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func InsertAjay(db *sql.DB, ajay User) (User, error) {
	stmt, err := db.Prepare("INSERT INTO ajay (name, age) VALUES (?, ?)")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(ajay.Name, ajay.Age)
	if err != nil {
		return User{}, err
	}

	// Retrieve the inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	ajay.ID = id

	return ajay, nil
}

func UpdateAjay(db *sql.DB, ajay *User) (*User, error) {
	stmt, err := db.Prepare("UPDATE ajay SET name = ?, age = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ajay.Name, ajay.Age, ajay.ID)
	if err != nil {
		return nil, err
	}

	// Fetch the updated record
	updatedAjay, err := GetAjay(db, ajay.ID)
	if err != nil {
		return nil, err
	}

	return &updatedAjay, nil
}
