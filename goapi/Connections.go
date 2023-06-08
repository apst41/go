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

type Ajay struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetAjay(db *sql.DB, id int64) (Ajay, error) {
	row := db.QueryRow("SELECT id, name, age FROM ajay WHERE id = ?", id)
	var ajay Ajay
	err := row.Scan(&ajay.ID, &ajay.Name, &ajay.Age)
	if err != nil {
		return Ajay{}, err
	}
	return ajay, nil
}

func FetchAllRows(db *sql.DB) ([]Ajay, error) {
	rows, err := db.Query("SELECT id, name, age FROM ajay")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Ajay

	for rows.Next() {
		var ajay Ajay
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
