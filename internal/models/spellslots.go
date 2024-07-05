package models

import "fmt"

type SpellSlot struct {
	Max       int `yaml:"max"`
	Remaining int `yaml:"remaining"`
	Level     int `yaml:"level"`
}

// Use a spell slot. Will return false if there are no remaining slots.
func (s *SpellSlot) Use() bool {
	if s.Remaining > 0 {
		s.Remaining--
		return true
	}
	return false
}

func (s *SpellSlot) Reset() {
	s.Remaining = s.Max
}

func (s *SpellSlot) String() string {
	return fmt.Sprintf("Level %d: %d/%d", s.Level, s.Remaining, s.Max)
}
