package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gopheramit/Learning-Python/pkg/models/mysql"
)

type application struct {
	logger *log.Logger
	users  *mysql.UserModel
	//logger *log.Logger
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", ":4000", "HTTP network address")
	// connStr := "user=postgres dbname=demo password=achal1234 host=localhost sslmode=disable"
	//dsn := flag.String("dsn", "user=postgres dbname=demo password=achal1234 host=localhost sslmode=disable", "PSQL data source name")

	dsn := flag.String("dsn", "web:achal@1234@/demo?parseTime=true", "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		logger: logger,
		users:  &mysql.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: logger,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()

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
