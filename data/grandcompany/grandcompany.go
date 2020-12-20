package grandcompany

// GrandCompany is the native representation of a Grand Company.
type GrandCompany uint8

const (
	None GrandCompany = iota
	Maelstrom
	OrderoftheTwinAdder
	ImmortalFlames
)

// Parse converts a string into the native representation of a Grand Company.
func Parse(input string) GrandCompany {
	switch input {
	case "Maelstrom":
		return Maelstrom
	case "Order of the Twin Adder":
		return OrderoftheTwinAdder
	case "Immortal Flames":
		return ImmortalFlames
	}
	return None
}
