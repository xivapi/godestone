package role

// Role is one of the five roles in FFXIV.
type Role string

// Role type.
const (
	None     Role = "None"
	Tank     Role = "Tank"
	Healer   Role = "Healer"
	DPS      Role = "DPS"
	Crafter  Role = "Crafter"
	Gatherer Role = "Gatherer"
)

// Parse converts the string representation of a role to its primitive equivalent.
func Parse(input string) Role {
	switch input {
	case "None":
		return None
	case "Tank":
		return Tank
	case "Healer":
		return Healer
	case "DPS":
		return DPS
	case "Crafter":
		return Crafter
	case "Gatherer":
		return Gatherer
	}
	return None
}
