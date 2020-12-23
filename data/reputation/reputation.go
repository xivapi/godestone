package reputation

// Reputation represents the alignment of an entity with a particular group.
type Reputation uint8

// Beast tribe or Grand Company reputation.
const (
	None Reputation = iota
	Neutral
	Recognized
	Friendly
	Trusted
	Respected
	Honored
	Sworn
	Allied
)

// Parse converts a string representation of a reputation to its primitive equivalent.
func Parse(input string) Reputation {
	switch input {
	case "Neutral":
		return Neutral
	case "Recognized":
		return Recognized
	case "Friendly":
		return Friendly
	case "Trusted":
		return Trusted
	case "Respected":
		return Respected
	case "Honored":
		return Honored
	case "Sworn":
		return Sworn
	case "Allied":
		return Allied
	}
	return None
}
