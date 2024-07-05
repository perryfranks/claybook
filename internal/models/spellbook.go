package models

import (
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type Spellbook struct {
	// I don't think there is a reason to pre-sort these. No need for performance
	Spells     []Spell     `yaml:"spells"`
	SpellSlots []SpellSlot `yaml:"spellslots"`
}

// this may end up with checks later on
func (sb *Spellbook) AddSpell(spell Spell) {
	sb.Spells = append(sb.Spells, spell)
}

type spellSlotByLevel []SpellSlot

func (s spellSlotByLevel) Len() int {
	return len(s)
}
func (s spellSlotByLevel) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s spellSlotByLevel) Less(i, j int) bool {
	return s[i].Level < s[j].Level
}

// sort the spell slots by level
func (sb *Spellbook) SortSpellSlots() {

	sort.Sort(spellSlotByLevel(sb.SpellSlots))

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
