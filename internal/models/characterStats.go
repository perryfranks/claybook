package models

import (
	"os"

	"gopkg.in/yaml.v3"
)

type CharacterStats struct {
	HitDiceSet HitDiceSet `yaml:"hitdice"`
}

func (cs *CharacterStats) Save(path string) error {

	data, err := yaml.Marshal(&cs)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (cs *CharacterStats) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// TODO: I wonder if this really works &cs
	if err := yaml.Unmarshal(data, &cs); err != nil {
		return err
	}

	return nil
}
