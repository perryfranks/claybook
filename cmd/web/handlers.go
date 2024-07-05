package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// Not needed with httprouter
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) spells(w http.ResponseWriter, r *http.Request) {

	app.spellbook.SortSpellSlots()
	data := app.newTemplateData(r)

	fmt.Println(data.String())

	app.render(w, http.StatusOK, "spellbook.tmpl", data)

}

func (app *application) updateSpellSlot(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("edit page"))
}

func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("edit page"))
}
