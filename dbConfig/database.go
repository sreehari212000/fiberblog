package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if errr := db.Ping(); errr != nil {
		return nil, errr
	}
	DB = db
	fmt.Println("CONNECTED TO DB")
	return DB, nil
}

func InsertIntoDb(name string, email string, password string) (sql.Result, error) {
	query := `INSERT INTO users (name, email, password) VALUES($1, $2, $3);`
	res, err := DB.Exec(query, name, email, password)
	if err != nil {
		log.Print("error inserting data into DB ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
