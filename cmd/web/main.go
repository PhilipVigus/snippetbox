package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:password@tcp(db:3306)/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var db *sql.DB // Declare db variable outside the loop

	for i := 0; i < 50; i++ {
		var err error
		db, err = openDB(*dsn)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Println("Failed to connect to database, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if db != nil {
		defer func() {
			err := db.Close()
			if err != nil {
				log.Println("Error closing database connection:", err)
			}
		}()
	} else {
		log.Fatal("Unable to connect to the database after multiple attempts")
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// create a pointer to a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
