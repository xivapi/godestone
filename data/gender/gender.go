package gender

// Gender is the gender of a character
type Gender uint8

const (
	None Gender = iota
	Male
	Female
)

// Parse returns the primitive representation of the provided gender.
func Parse(input string) Gender {
	switch input {
	case "♂":
		return Male
	case "♀":
		return Female

	}
	return None
}
