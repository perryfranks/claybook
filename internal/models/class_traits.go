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

func ClassName(class Class) string {
	return class.String()
}

const (
	Pugilist Class = iota
	Sorcerer
)

type ClassTraitsSet struct {
	ClassTraits map[Class]ClassTraits `yaml:"ClassTraits"`
}

// NOP on duplicate
// Need to Add what to do if the map is empty
func (cts *ClassTraitsSet) Add(ct ClassTraits) {

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

// Render the ClassTraitsSet as a list for display
func (cts *ClassTraitsSet) List() []ClassTraits {
	var l []ClassTraits
	for _, values := range cts.ClassTraits {
		l = append(l, values)
	}

	return l

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
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}
