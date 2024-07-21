package models

import "math"

type AbilityScores struct {
	Str int `yaml:"Strength"`
	Dex int `yaml:"Dexterity"`
	Con int `yaml:"Constitution"`
	Int int `yaml:"Intelligence"`
	Wis int `yaml:"Wisdom"`
	Cha int `yaml:"Charisma"`
}

func AbilityScoreMod(score int) int {
	fres := float64(score-10) / float64(2)
	return int(math.Floor(fres))
}
