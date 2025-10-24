package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createTables(db *sql.DB) {
	fmt.Println("CREATING USER TABLE...")
	queryString := `CREATE TABLE IF NOT EXISTS users (
		user_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS posts (
		post_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		author_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS comments (
		comment_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		post_id BIGINT NOT NULL REFERENCES posts(post_id) ON DELETE CASCADE,
		author_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
		text TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS likes (
		like_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
		post_id BIGINT NOT NULL REFERENCES posts(post_id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, post_id)
	);`
	_, err := db.Exec(queryString)
	if err != nil {
		fmt.Println("ERROR CREATING TABLES ", err)
		return
	}
	fmt.Println("Tables created successfully!")
}

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal(envError)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createTables(db)
}
