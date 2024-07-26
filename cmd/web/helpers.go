package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"claybook.perryfranks.nerd/internal/models"
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

// try and execute the block level template not the base
// it would be possible to get the name of the template and then take that as the template name
// but that would be an invisible constraint
func (app *application) renderBlock(w http.ResponseWriter, status int, page string, templateName string, data *templateData) {

	// get template from cache if it exists
	// ts, ok := app.templateCache[page]
	ts, ok := app.partialsCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, templateName, data)
	// err := ts.Execute(buf, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	fmt.Println(buf)
	buf.WriteTo(w)

}

// Create a new data struct for any common data that we don't mind passing to all templates
func (app *application) newTemplateData(r *http.Request) *templateData {
	sbl := app.spellbook.GetSortedSpellsByLevel()
	ctl := app.characterStats.ClassTraitsSet.List()
	return &templateData{
		Spellbook:         &app.spellbook.Spells,
		SpellSlots:        &app.spellbook.SpellSlots,
		SpellsByLevel:     &sbl,
		HitDiceSet:        &app.characterStats.HitDiceSet.HitDice,
		MoxiePoints:       &app.characterStats.MoxiePoints,
		AbilityScores:     &app.characterStats.AbilityScores,
		ClassTraitsList:   &ctl,
		Features:          &app.characterStats.Features.Feats,
		CombatStats:       &app.characterStats.CombatStats,
		Languages:         &app.characterStats.Languages,
		ToolProficiencies: &app.characterStats.ToolProficiencies,
		PugilistDie:       &app.characterStats.PugilistDie,
		Attacks:           &app.characterStats.Attacks,
		ProficiencyBonus:  &app.characterStats.ProficiencyBonus,
		Inventory:         &app.characterStats.Inventory,
		MiscItems:         &app.miscItems,
	}
}

// Save and load. Load will panic on error
func (app *application) saveData() {
	app.characterStats.Save(app.savefiles["characterStats"])
	models.SaveSpellbook(app.savefiles["spells"], app.spellbook)
}

// Save and load. Load will panic on error
func (app *application) loadData() {
	var err error
	app.spellbook, err = models.LoadSpellbook(app.savefiles["spells"])
	if err != nil {
		panic(err)
	}

	err = app.characterStats.Load(app.savefiles["characterStats"])
	if err != nil {
		panic(err)
	}

}
