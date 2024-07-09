package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// 404 convience function that just fits with the others
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {

	// get template from cache if it exists
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)

}

// Create a new data struct for any common data that we don't mind passing to all templates
func (app *application) newTemplateData(r *http.Request) *templateData {
	sbl := app.spellbook.GetSortedSpellsByLevel()
	return &templateData{
		Spellbook:     &app.spellbook.Spells,
		SpellSlots:    &app.spellbook.SpellSlots,
		SpellsByLevel: &sbl,
		HitDiceSet:    &app.characterStats.HitDiceSet.HitDice,
	}

}
