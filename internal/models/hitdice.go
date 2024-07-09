package models

type HitDice struct {
	Number int `yaml:"number"`
	Sides  int `yaml:"sides"`
}

type HitDiceSet struct {
	HitDice []HitDice `yaml:"hitdice"`
}

// If a number of sides match then add the number of the dice. Otherwise, the hitdice is appended to the set
func (h *HitDiceSet) add(hd HitDice) {
	for i, existing := range h.HitDice {
		if existing.Sides == hd.Sides {
			h.HitDice[i].Number += hd.Number
		} else {
			h.HitDice = append(h.HitDice, hd)
		}
	}
}
