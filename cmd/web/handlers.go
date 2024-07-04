package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/character.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Println(err.Error())
		// // http.Error(w, "Internal Server error", 500)
		//
		// return
		app.serverError(w, err)
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}
}

func (app *application) spells(w http.ResponseWriter, r *http.Request) {
	// printout := fmt.Sprintf("%v", app.spellbook)
	// w.Write([]byte(printout))

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/spellbook.tmpl",
		"./ui/html/partials/spellcard.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		Spellbook: &app.spellbook.Spells,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("edit page"))
}
