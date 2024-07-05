package models

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
