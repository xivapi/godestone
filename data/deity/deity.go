package deity

// GuardianDeity is one of the twelve guardian deities a character can be associated with.
type GuardianDeity uint8

// FFXIV guardian deity.
const (
	None GuardianDeity = iota
	HalonetheFury
	MenphinatheLover
	ThaliaktheScholar
	NymeiatheSpinner
	LlymlaentheNavigator
	OschontheWanderer
	ByregottheBuilder
	RhalgrtheDestroyer
	AzeymatheWarden
	NaldthaltheTraders
	NophicatheMatron
	AlthyktheKeeper
)

// Parse converts a guardian deity name into its native representation.
func Parse(input string) GuardianDeity {
	switch input {
	case "Halone, the Fury":
		return HalonetheFury
	case "Menphina, the Lover":
		return MenphinatheLover
	case "Thaliak, the Scholar":
		return ThaliaktheScholar
	case "Nymeia, the Spinner":
		return NymeiatheSpinner
	case "Llymlaen, the Navigator":
		return LlymlaentheNavigator
	case "Oschon, the Wanderer":
		return OschontheWanderer
	case "Byregot, the Builder":
		return ByregottheBuilder
	case "Rhalgr, the Destroyer":
		return RhalgrtheDestroyer
	case "Azeyma, the Warden":
		return AzeymatheWarden
	case "Nald'thal, the Traders":
		return NaldthaltheTraders
	case "Nophica, the Matron":
		return NophicatheMatron
	case "Althyk, the Keeper":
		return AlthyktheKeeper
	}
	return None
}
