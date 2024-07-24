package main

import (
	"html/template"
	"path/filepath"

	"claybook.perryfranks.nerd/internal/models"
)

// lookup for names of functions and their functions
var functions = template.FuncMap{
	"Mod":       models.AbilityScoreMod,
	"ClassName": models.ClassName,
}

// holding structure for any data we want to pass to our
type templateData struct {
	Spellbook         *[]models.Spell
	Spell             *models.Spell // So far unused
	SpellSlots        *[]models.SpellSlot
	SpellsByLevel     *[]models.SpellByLevel
	HitDiceSet        *[]models.HitDice
	MoxiePoints       *models.MoxiePoints
	AbilityScores     *models.AbilityScores
	ClassTraitsList   *[]models.ClassTraits
	Features          *[]models.Trait
	CombatStats       *models.CombatStats
	Languages         *[]string
	ToolProficiencies *[]string
	PugilistDie       *int
	Attacks           *[]models.Attack
}

func (t *templateData) String() string {
	printout := ""
	if t.Spellbook != nil {
		printout += "Spellbook: "
		for _, spell := range *t.Spellbook {
			printout += spell.Name + ", "
		}
	} else {
		printout += "Spellbook: nil"
	}

	if t.SpellSlots != nil {
		printout += "\nSpellSlots: "
		for _, slot := range *t.SpellSlots {
			printout += slot.String() + ", "
		}
	} else {
		printout += "\nSpellSlots: nil"
	}

	if t.Spell != nil {
		printout += "\nSpell: " + t.Spell.Name
	} else {
		printout += "\nSpell: nil"
	}

	return printout

}

func newTemplateData() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract the file name (like "home.tmpl") from the full path
		name := filepath.Base(page)

		// Register the template functions as well.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")

		// ts, err := template.ParseFiles("./ui/html/base.tmpl", page)
		if err != nil {
			return nil, err
		}

		// grab all the partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
