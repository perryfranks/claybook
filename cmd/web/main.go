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

	// Character stuff
	characterStats models.CharacterStats
	spellbook      models.Spellbook
	savefiles      map[string]string
}

func main() {

	// _ = testSpells()

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// template cache
	templateCache, err := newTemplateData()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,

		spellbook:      models.Spellbook{},
		characterStats: models.CharacterStats{},
		savefiles: map[string]string{
			"spells":         "internal/data/spells.yaml",
			"characterStats": "internal/data/characterStats.yaml",
		},
	}

	app.spellbook, err = models.LoadSpellbook(app.savefiles["spells"])
	if err != nil {
		panic(err)
	}

	err = app.characterStats.Load(app.savefiles["characterStats"])
	if err != nil {
		panic(err)
	}

	// NOTE: these are actually never called how I run the sever
	defer app.characterStats.Save(app.savefiles["characterStats"])
	defer models.SaveSpellbook(app.savefiles["spells"], app.spellbook)

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
