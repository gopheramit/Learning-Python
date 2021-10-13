package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gopheramit/Learning-Python/pkg/models/psql"

	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	users    *psql.UserModel
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", ":4000", "HTTP network address")
	// connStr := "user=postgres dbname=demo password=achal1234 host=localhost sslmode=disable"
	//dsn := flag.String("dsn", "user=postgres dbname=demo password=achal1234 host=localhost sslmode=disable", "PSQL data source name")

	dsn := flag.String("dsn", "user=postgres1 dbname=demo password=achal1234 host=localhost sslmode=disable", "PSQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
		log.Println(err)
	}
	defer db.Close()
	flag.Parse()

	app := &application{
		errorLog: errorLog,
		users:    &psql.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()

	errorLog.Fatal(err)

}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
		panic(err)

	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return db, nil
}
