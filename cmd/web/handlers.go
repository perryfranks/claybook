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

func (app *application) useHitDice(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	rawSides := r.PostForm.Get("hit-dice")
	sides, err := strconv.Atoi(rawSides)
	if err == nil {
		// find the dice in the character's hit dice
		for i, hd := range app.characterStats.HitDiceSet.HitDice {
			if hd.Sides == sides {
				app.characterStats.HitDiceSet.HitDice[i].Use()
			}

		}
	}

	// check for reset
	if r.PostForm.Get("reset") == "true" {
		fmt.Println("Resetting Hit Dice")
		app.characterStats.HitDiceSet.Reset()
	}

	// save the character stats
	app.characterStats.Save(app.savefiles["characterStats"])

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) useMoxie(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	use := r.PostForm.Get("use")
	if use == "true" {
		app.characterStats.MoxiePoints.Use()
	}

	app.characterStats.Save(app.savefiles["characterStats"])

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) resetMoxie(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	use := r.PostForm.Get("reset")
	if use == "true" {
		app.characterStats.MoxiePoints.Reset()
	}

	app.characterStats.Save(app.savefiles["characterStats"])

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) changeHealth(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	changeRaw := r.PostForm.Get("healthChange")
	change, err := strconv.Atoi(changeRaw)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.characterStats.CombatStats.ChangeHealth(change)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) save(w http.ResponseWriter, r *http.Request) {
	app.characterStats.Save(app.savefiles["characterStats"])
	models.SaveSpellbook(app.savefiles["spells"], app.spellbook)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) load(w http.ResponseWriter, r *http.Request) {
	err := app.characterStats.Load(app.savefiles["characterStats"])
	if err != nil {
		app.errorLog.Println(err)
	}

	// err = app.spellbook.Load(app.savefiles["spells"])
	app.spellbook, err = models.LoadSpellbook(app.savefiles["spells"])
	if err != nil {
		app.errorLog.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) classTraits(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Comming soon")
}
