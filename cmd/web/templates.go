package main

import (
	"html/template"
	"path/filepath"

	"claybook.perryfranks.nerd/internal/models"
)

// lookup for names of functions and their functions
var functions = template.FuncMap{
	"Mod":         models.AbilityScoreMod,
	"ClassName":   models.ClassName,
	"AttackBonus": models.AttackBonus,
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
	ProficiencyBonus  *int
	Inventory         *models.Inventory
	MiscItems         *[]models.Item
	DataDump          *[]string
	FightClubFeatures *[]models.Trait
	FightClubName     *string
	CharacterDump     *string
	SpellsDump        *string
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

// tmpl -> html
func newTemplateData() (map[string]*template.Template, map[string]*template.Template, error) {

	cache := map[string]*template.Template{}
	partialsCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, nil, err
	}

	for _, page := range pages {
		// extract the file name (like "home.html") from the full path
		name := filepath.Base(page)

		// Register the template functions as well.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")

		// ts, err := template.ParseFiles("./ui/html/base.html", page)
		if err != nil {
			return nil, nil, err
		}

		// grab all the partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, nil, err
		}

		cache[name] = ts

	}

	// With the inclusion of htmx we now want to be able to have the partials inclueded as well. Not just the compiled top-level pages
	partials, err := filepath.Glob("./ui/html/partials/*.html")
	if err != nil {
		return nil, nil, err
	}

	// Get the funcs back in as well
	for _, partial := range partials {
		name := filepath.Base(partial)

		ts, err := template.New(name).Funcs(functions).ParseFiles(partial)
		if err != nil {
			return nil, nil, nil
		}

		partialsCache[name] = ts
	}

	return cache, partialsCache, nil
}
