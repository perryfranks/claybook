package main

import (
	"claybook.perryfranks.nerd/internal/models"
)

// holding structure for any data we want to pass to our
type templateData struct {
	Spellbook  *[]models.Spell
	Spell      *models.Spell // So far unused
	SpellSlots *[]models.SpellSlot
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
