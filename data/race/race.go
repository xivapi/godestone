package race

// Race is the general race of a character.
type Race uint8

const (
	None Race = iota
	Hyur
	Elezen
	Lalafell
	Miqote
	Roegadyn
	AuRa
	Hrothgar
	Viera
)

// Parse returns the primitive representation of the provided race.
func Parse(input string) Race {
	switch input {
	case "Hyur":
		return Hyur
	case "Elezen":
		return Elezen
	case "Lalafell":
		return Lalafell
	case "Miqo'te":
		return Miqote
	case "Roegadyn":
		return Roegadyn
	case "Au Ra":
		return AuRa
	case "Hrothgar":
		return Hrothgar
	case "Viera":
		return Viera
	}
	return None
}
