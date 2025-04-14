package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Author struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Experience int    `db:"experience" json:"experience"`
}

type Book struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	AuthorID int    `db:"author_id" json:"author_id"`
	Year     int    `db:"year" json:"year"`
}

func initDB() (*sqlx.DB, error) {
	connStr := "user=postgres password=1234 dbname=lab3 sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open error: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}
	return db, nil
}

func SearchAuthorBooks(db *sqlx.DB, a Author) ([]Book, bool, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE author_id = $1", a.ID)
	if err != nil {
		return nil, false, err
	}
	return books, len(books) > 0, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	defer db.Close()

	author := Author{ID: 1}
	books, found, err := SearchAuthorBooks(db, author)
	if err != nil {
		log.Fatalf("Search error: %v", err)
	}

	if found {
		fmt.Printf("Книги автора:\n")
		for _, b := range books {
			fmt.Printf("- %s (%d)\n", b.Title, b.Year)
		}
	} else {
		fmt.Println("Книг не знайдено.")
	}
}

// Added for lab4 PR

