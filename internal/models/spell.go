package models

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Spell struct {
	Level       int    `yaml:"level"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	HigherLevel string `yaml:"higherlevel"`
	RangeFeet   int    `yaml:"range"`
	Duration    string `yaml:"duration"`
	Casting     string `yaml:"casting"`
	Components  string `yaml:"components"`
	School      string `yaml:"school"`
	// could add a roll function with some thinking
}

type SpellSlot struct {
	max  int `yaml:"max"`
	used int `yaml:"used"`
}

type Spellbook struct {
	// I don't think there is a reason to pre-sort these. No need for performance
	Spells     []Spell           `yaml:"spells"`
	SpellSlots map[int]SpellSlot `yaml:"spellslots"`
}

// this may end up with checks later on
func (sb *Spellbook) AddSpell(spell Spell) {
	sb.Spells = append(sb.Spells, spell)
}

// Add a spell and then call the SaveSpellBook function on exit
func (sb *Spellbook) AddAndSave(spell Spell, file string) {
	sb.AddSpell(spell)
	SaveSpellbook(file, *sb)
}

// Probable Functions:
// Get spells for level x
// Read from file
func LoadSpellbook(file string) (Spellbook, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		// panic(err)
		return Spellbook{}, err
	}

	var spellbook Spellbook
	if err := yaml.Unmarshal(data, &spellbook); err != nil {
		return Spellbook{}, err
	}

	return spellbook, nil
}

// Save from file
func SaveSpellbook(file string, sb Spellbook) error {
	data, err := yaml.Marshal(&sb)
	if err != nil {
		// panic(err)
		return err
	}

	err = os.WriteFile(file, data, 0644)

	if err != nil {
		// panic(err)
		return err
	}

	return nil
}
