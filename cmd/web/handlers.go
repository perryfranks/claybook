package main

import (
	"fmt"
	"net/http"
	"strconv"

	"claybook.perryfranks.nerd/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// Not needed with httprouter
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "character.tmpl", data)
}

func (app *application) spells(w http.ResponseWriter, r *http.Request) {

	app.spellbook.SortSpellSlots()
	data := app.newTemplateData(r)

	fmt.Println(data.String())

	app.render(w, http.StatusOK, "spellbook.tmpl", data)

}

func (app *application) updateSpellSlot(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the slot level
	rawSlot := r.PostForm.Get("slotlevel")
	slotLevel, err := strconv.Atoi(rawSlot)
	if err == nil {
		app.infoLog.Println("Slot Level: ", slotLevel)
		dbSlots := app.spellbook.SpellSlots
		for i, slot := range dbSlots {
			if slot.Level == slotLevel {
				dbSlots[i].Use()
			}
		}

	}

	// check for reset button
	if r.PostForm.Get("reset") == "true" {
		app.infoLog.Println("Resetting")
		for i := range app.spellbook.SpellSlots {
			// *slot.Reset()
			app.spellbook.SpellSlots[i].Reset()
		}
	}

	// update the spell slot
	// get the corresponding spell slot

	// save
	// app.spellbook.("internal/data/spells.yaml")
	models.SaveSpellbook("internal/data/spells.yaml", app.spellbook)
	// http.Redirect(w, r, "/spells", http.StatusSeeOther)
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "spellbook.tmpl", data)

}

func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("edit page"))
}
