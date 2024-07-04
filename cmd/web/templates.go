package main

import "claybook.perryfranks.nerd/internal/models"

// holding structure for any data we want to pass to our
type templateData struct {
	Spellbook *[]models.Spell
	Spell     *models.Spell // So far unused
}
