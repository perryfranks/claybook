package models

type Trait struct {
	Level       int    `yaml:"level"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// Just all the little character features and what not that you get
// 0 can be innate
type Features struct {
	Feats []Trait `yaml:"feats"`
}
