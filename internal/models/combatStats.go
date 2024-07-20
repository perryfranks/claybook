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