package models

// Try and use a go ennum
type Class int64

func (c *Class) String() string {
	switch *c {
	case Pugilist:
		return "Pugilist"
	case Sorcerer:
		return "Sorcrerer"
	}
	return ""
}

const (
	Pugilist Class = iota
	Sorcerer
)

type ClassTraitsSet struct {
	ClassTraits map[Class]ClassTraits `yaml:"ClassTraits"`
}

// NOP on duplicate
// Need to add what to do if the map is empty
func (cts *ClassTraitsSet) add(ct ClassTraits) {

	// create the map if it's empty
	if cts.ClassTraits == nil {
		cts.ClassTraits = make(map[Class]ClassTraits)

	}

	_, ok := cts.ClassTraits[ct.ClassName]
	if ok == true {
		// NOP
	} else {
		// new value check for ennum
		switch ct.ClassName {
		case Pugilist:
			cts.ClassTraits[ct.ClassName] = ct
		case Sorcerer:
			cts.ClassTraits[ct.ClassName] = ct
		}
	}
}

type ClassTraits struct {
	ClassName Class   `yaml:"Class"`
	Traits    []Trait `yaml:"Traits"`
}

func (ct *ClassTraits) add(t Trait) {
	ct.Traits = append(ct.Traits, t)
}

type Trait struct {
	Level       int    `yaml:"level"`
	Description string `yaml:"description"`
}
