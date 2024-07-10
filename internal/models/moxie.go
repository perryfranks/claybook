package models

type MoxiePoints struct {
	Current int `yaml:"current"`
	Max     int `yaml:"max"`
}

func (mp *MoxiePoints) Use() {
	mp.Current--
}

func (mp *MoxiePoints) Reset() {
	mp.Current = mp.Max
}
