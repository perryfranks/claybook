package models

// Init everything with dummy data so I don't need to think about how the yaml should look
// If it happens to be the current and correct character sheet oh well
func (cs *CharacterStats) SetFakeData() {
	abilityscores := AbilityScores{
		Str: 15,
		Dex: 8,
		Con: 16,
		Int: 10,
		Wis: 12,
		Cha: 13,
	}

	hitDiceSet := HitDiceSet{
		HitDice: []HitDice{
			{
				Number:  3,
				Current: 3,
				Sides:   8,
			},
			{
				Number:  1,
				Current: 1,
				Sides:   6,
			},
		},
	}

	moxie := MoxiePoints{
		Current: 2,
		Max:     2,
	}

	combatStats := CombatStats{
		HpMax:      37,
		HpCurrent:  37,
		HpTemp:     0,
		ArmorClass: 15,
		DeathSaves: 0,
	}

	languages := []string{"common", "Yaun-ti Lizard", "Mystery"}
	toolProficiencies := []string{"improvised weapons", "whip", "crossbow", "tool set", "light armour", "simple weapons"}

	classTraitsSet := ClassTraitsSet{
		ClassTraits: map[Class]ClassTraits{
			Pugilist: {
				ClassName: Pugilist,
				Traits: []Trait{
					{
						Level:       1,
						Name:        "Iron Chin",
						Description: "Beginning at 1st level, while you are wearing light or no armor and not wielding a shield, your AC equals 12 + your Constitution modifier.",
					},
					{
						Level:       1,
						Name:        "Fisticuffs",
						Description: "At 1st level, your years of fighting in back alleys and taverns have given you mastery over combat styles that use unarmed strikes and pugilist weapons, which are simple melee weapons without the two-handed property, whips, and improvised weapons. You can’t use the finesse property of a weapon while using it as a pugilist weapon. \nYou gain the following benefits while you are unarmed or using only pugilist weapons and you are wearing light or no armor and not using a shield: \n» You can roll a d6 in place of the normal damage of your unarmed strike or pugilist weapon. This die changes as you gain pugilist levels, as shown in the Fisticuffs column on the Pugilist table. \n» When you use the Attack action on your turn and make only unarmed strikes, attacks with pugilist weapons, shoves, or grapples, you can use a bonus action to make one grapple or unarmed strike",
					},
					{
						Level:       2,
						Name:        "Moxie",
						Description: "Your experience laying the beatdown on others has given you a moxie you can channel in the midst of battle. This swagger is represented by a number of moxie points. Your pugilist level determines the maximum number of points you have, as shown in the Moxie Points column of the Pugilist table. \nYou can spend these points to fuel various moxie features. You start knowing three such features: Brace Up, The Old One-Two, and Stick and Move. You regain all expended moxie points when you finish a short or long rest.",
					},
					{
						Level:       2,
						Name:        "Brace Up",
						Description: "You can use a bonus action and spend 1 moxie point to brace for attacks. You gain a number of temporary hit points equal to a roll of your fisticuffs die + your pugilist level + your Constitution modifier. You lose any remaining temporary hit points gained in this way after 1 minute.",
					},
					{
						Level:       2,
						Name:        "The Old One-Two",
						Description: "Immediately after you take the Attack action on your turn, you can spend 1 moxie point to make two unarmed strikes as a bonus action",
					},
					{
						Level:       2,
						Name:        "Stick and Move",
						Description: "You can use a bonus action and expend 1 moxie point to attempt to shove a creature or take the Dash action.",
					},
					{
						Level:       2,
						Name:        "Street Smart",
						Description: "Carousing, shadowboxing, and sparring all count as light activity for the purposes of resting for you. Additionally, once you have caroused in a settlement for 8 hours or more, you know all public locations in the city as if you were born and raised there and you cannot be lost by nonmagical means while within the city.",
					},
					{
						Level:       3,
						Name:        "Bloodied But Unbowed",
						Description: "When you take damage that reduces you to half your maximum hit points or less, you can use your reaction to gain temporary hit points equal to three times your pugilist level and regain all your expended moxie points. Once you use this feature, you can’t use it again until you finish a short or long rest",
					},
					{
						Level:       3,
						Name:        "Fight Club",
						Description: "You get a fight Club",
					},
				},
			},
			Sorcerer: {
				ClassName: Sorcerer,
				Traits: []Trait{
					{
						Level:       1,
						Name:        "Dragon Ancestor",
						Description: "At 1st level, you choose one type of dragon as your ancestor. The damage type associated with each dragon is used by features you gain later. You can speak, read, and write Draconic. Additionally, whenever you make a Charisma check when interacting with dragons, your proficiency bonus is doubled if it applies to the check.\nType: Copper: Acid",
					},
				},
			},
		},
	}

	// Now lets set everthing
	cs.AbilityScores = abilityscores
	cs.HitDiceSet = hitDiceSet
	cs.MoxiePoints = moxie
	cs.CombatStats = combatStats
	cs.Languages = languages
	cs.ToolProficiencies = toolProficiencies
	cs.ClassTraitsSet = classTraitsSet

}