package town

// Town is one of the major cities.
type Town uint8

const (
	Nowheresville Town = iota
	LimsaLominsa
	Gridania
	Uldah
	Ishgard
	Kugane        Town = 7
	TheCrystarium Town = 10
)

// Parse returns the primitive representation of the provided town.
func Parse(input string) Town {
	switch input {
	case "Limsa Lominsa":
		return LimsaLominsa
	case "Gridania":
		return Gridania
	case "Ul'dah":
		return Uldah
	case "Ishgard":
		return Ishgard
	case "Kugane":
		return Kugane
	case "The Crystarium":
		return TheCrystarium
	}
	return Nowheresville
}
