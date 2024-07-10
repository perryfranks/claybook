package models

type HitDice struct {
	Number  int `yaml:"number"`
	Current int `yaml:"current"`
	Sides   int `yaml:"sides"`
}

func (h *HitDice) Use() {
	h.Current--
}

type HitDiceSet struct {
	HitDice []HitDice `yaml:"hitdice"`
}

// If a number of sides match then add the number of the dice. Otherwise, the hitdice is appended to the set
func (hds *HitDiceSet) add(hd HitDice) {
	for i, existing := range hds.HitDice {
		if existing.Sides == hd.Sides {
			hds.HitDice[i].Number += hd.Number
		} else {
			hds.HitDice = append(hds.HitDice, hd)
		}
	}
}

func (hds *HitDiceSet) Reset() {
	for i := range hds.HitDice {
		hds.HitDice[i].Current = hds.HitDice[i].Number
	}
}
