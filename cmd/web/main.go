package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"claybook.perryfranks.nerd/
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func testSpells() SpellBook {

}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server sever on ", *addr)
	err := srv.ListenAndServe()
	// err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
