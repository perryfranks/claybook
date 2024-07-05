package main

import (
	"fmt"
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
		app.serverError(w, err)
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) spells(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/spellbook.tmpl",
		"./ui/html/partials/spellcard.tmpl",
		"./ui/html/partials/spellslotbar.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.spellbook.SortSpellSlots()
	data := &templateData{
		Spellbook: &app.spellbook.Spells,
		// TODO: will probs need to sort this
		SpellSlots: &app.spellbook.SpellSlots,
	}

	fmt.Println(data.String())

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("edit page"))
}
