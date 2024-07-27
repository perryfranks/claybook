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

	app.saveData()
	data := app.newTemplateData(r)
	app.renderBlock(w, http.StatusOK, "spellslotbar.tmpl", "spellslotcard", data)

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

	data := app.newTemplateData(r)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	app.renderBlock(w, http.StatusOK, "hitdice.tmpl", "hit-dice", data)

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

	data := app.newTemplateData(r)
	app.renderBlock(w, http.StatusOK, "moxie.tmpl", "moxie", data)
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

	data := app.newTemplateData(r)
	app.renderBlock(w, http.StatusOK, "moxie.tmpl", "moxie", data)

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
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "classTraits.tmpl", data)
	return
}

func (app *application) features(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "features.tmpl", data)
	return
}

func (app *application) updateHpDamage(w http.ResponseWriter, r *http.Request) {

	// Now we need to get the data from the form. Validate that it's a number and then go from there
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the damage
	rawDamage := r.PostForm.Get("damage")
	damage, err := strconv.Atoi(rawDamage)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.characterStats.CombatStats.ChangeHealth(damage)

	app.saveData()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) updateHpHeal(w http.ResponseWriter, r *http.Request) {
	// Now we need to get the data from the form. Validate that it's a number and then go from there
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the damage
	rawHeal := r.PostForm.Get("heal")
	heal, err := strconv.Atoi(rawHeal)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.characterStats.CombatStats.Heal(heal)

	app.saveData()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) updateHpTemp(w http.ResponseWriter, r *http.Request) {
	// Now we need to get the data from the form. Validate that it's a number and then go from there
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the amount
	rawTemp := r.PostForm.Get("temp")
	temp, err := strconv.Atoi(rawTemp)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fmt.Println("temp amount ", temp)
	app.characterStats.CombatStats.Temp(temp)

	app.saveData()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) misc(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "misc.tmpl", data)
}

func (app *application) inventory(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)
	fmt.Println("inventory called")
	app.render(w, http.StatusOK, "inventory.tmpl", data)
}

func (app *application) inventoryAdd(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	cost, err := strconv.Atoi(r.PostForm.Get("cost"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	weight := r.PostForm.Get("weight")
	description := r.PostForm.Get("description")

	i := models.Item{
		Name:        name,
		Cost:        cost,
		Weight:      weight,
		Description: description,
	}

	app.characterStats.Inventory.Add(i)
	app.saveData()

	data := app.newTemplateData(r)
	// http.Redirect(w, r, "/inventory", http.StatusSeeOther)
	app.renderBlock(w, http.StatusOK, "inventorylist.tmpl", "inventory-list", data)

}

// dump all character data nothing fancy
func (app *application) dumpCharacter(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "dataDump.tmpl", data)

}
