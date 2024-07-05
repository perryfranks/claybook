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
	spellbook     models.Spellbook
	templateCache map[string]*template.Template
}

func testSpells() models.Spellbook {
	sb := models.Spellbook{}
	sb.AddSpell(models.Spell{})
	models.SaveSpellbook("internal/data/spells.yaml", sb)

	return sb
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

	sb, err := models.LoadSpellbook("internal/data/spells.yaml")
	if err != nil {
		panic(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		spellbook:     sb,
		templateCache: templateCache,
	}
	defer models.SaveSpellbook("internal/data/spells.yaml", app.spellbook)

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
