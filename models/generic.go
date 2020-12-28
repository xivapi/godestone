package models

// NamedEntity represents an instance of an object with a name.
type NamedEntity struct {
	ID   uint32
	Name string
	Icon string

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}

// GenderedEntity represents an instance of an object with masculine and feminine names.
type GenderedEntity struct {
	ID   uint32
	Name string

	NameMasculineEN string
	NameMasculineJA string
	NameMasculineDE string
	NameMasculineFR string
	NameFeminineEN  string
	NameFeminineJA  string
	NameFeminineDE  string
	NameFeminineFR  string
}
