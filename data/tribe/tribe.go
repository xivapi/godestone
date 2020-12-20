package tribe

// Tribe is the specific race of a character.
type Tribe = uint8

const (
	None Tribe = iota
	Midlander
	Highlander
	Wildwood
	Duskwight
	Plainsfolk
	Dunesfolk
	SeekeroftheSun
	KeeperoftheMoon
	SeaWolf
	Hellsguard
	Raen
	Xaela
	Helions
	TheLost
	Rava
	Veena
)

// Parse returns the primitive representation of the provided tribe.
func Parse(input string) Tribe {
	switch input {
	case "Midlander":
		return Midlander
	case "Highlander":
		return Highlander
	case "Wildwood":
		return Wildwood
	case "Duskwight":
		return Duskwight
	case "Plainsfolk":
		return Plainsfolk
	case "Dunesfolk":
		return Dunesfolk
	case "Seeker of the Sun":
		return SeekeroftheSun
	case "Keeper of the Moon":
		return KeeperoftheMoon
	case "Sea Wolf":
		return SeaWolf
	case "Hellsguard":
		return Hellsguard
	case "Raen":
		return Raen
	case "Xaela":
		return Xaela
	case "Helions":
		return Helions
	case "The Lost":
		return TheLost
	case "Rava":
		return Rava
	case "Veena":
		return Veena
	}
	return None
}
