package models

type Spell struct {
	level       int
	name        string
	description string
	rangeFeet   int
	duration    string
	casting     string
	components  string
	school      string

	// could add a roll function with some thinking
}
