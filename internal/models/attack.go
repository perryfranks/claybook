package models

type Attack struct {
	Name  string `yaml:"name"`
	Melee bool   `yaml:"isMelee"`
	Dice  string `yaml:"dice"`
}

// TODO: need a way to calculate the damage bonus
