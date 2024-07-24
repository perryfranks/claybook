package models

// Container for stats mainly used in combat
type CombatStats struct {
	HpMax      int
	HpCurrent  int
	HpTemp     int
	ArmorClass int
	DeathSaves int
}

func (cs *CombatStats) ChangeHealth(change int) {
	// remove from temp first
	cs.HpTemp -= change
	if cs.HpTemp < 0 {
		cs.HpCurrent += cs.HpTemp
		cs.HpTemp = 0
	}

	if cs.HpCurrent < 0 {
		cs.HpCurrent = 0
	}
}

// updating the temp health. It would be good to sort out only allowing one source of temp hit points at a time
// Heal raise current until it is HpMax
func (cs *CombatStats) Heal(change int) {

	cs.HpCurrent += change
	if cs.HpCurrent > cs.HpMax {
		cs.HpCurrent = cs.HpMax
	}
}

func (cs *CombatStats) Temp(change int) {
	cs.HpTemp += change
}
