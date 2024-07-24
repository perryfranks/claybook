package models

type Attack struct {
	Name  string
	Melee bool
	Roll  Roll
}

type Roll struct {
}
