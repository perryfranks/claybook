package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"claybook.perryfranks.nerd/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
	// partialsCache map[string]*template.Template

	// Character stuff
	characterStats models.CharacterStats
	spellbook      models.Spellbook
	miscItems      []models.Item // Just random notes that things I've missed
	savefiles      map[string]string
	DisplayVars    DisplayVars
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// template cache
	templateCache, err := newTemplateData()
	if err != nil {
		errorLog.Fatal(err)
	}

	dv := DisplayVars{
		NavBarColors: map[string]string{},
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
		// partialsCache: partialsCache,

		spellbook:      models.Spellbook{},
		characterStats: models.CharacterStats{},
		savefiles: map[string]string{
			"spells":         "internal/data/spells.yaml",
			"characterStats": "internal/data/characterStats.yaml",
		},
		DisplayVars: dv,
	}

	app.loadData()

	// TODO: Remove this at some point
	app.characterStats.SetFakeData()

	// NOTE: these are actually never called how I run the sever
	defer app.saveData()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server sever on ", *addr)
	err = srv.ListenAndServe()
	// err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
