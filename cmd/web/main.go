package main

import "fmt"

type application struct{


}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	app := &application{
		
	}
	

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
}
