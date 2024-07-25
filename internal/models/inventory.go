package models

type Inventory struct {
	Items     []Item `yaml:"Items"`
	MaxWeight int    `yaml:"MaxWeight"`
}

type Item struct {
	Name        string `yaml:"Name"`
	Cost        int    `yaml:"Cost"`
	Weight      string `yaml:"Weight"`
	Description string `yaml:"Description"`
}

func (inv *Inventory) Add(i Item) {
	inv.Items = append(inv.Items, i)
}
