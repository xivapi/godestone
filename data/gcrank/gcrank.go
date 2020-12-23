package gcrank

// GCRank is the Grand Company rank of a character.
type GCRank uint8

// FFXIV Grand Company rank.
const (
	None GCRank = iota
	PrivateThirdClass
	PrivateSecondClass
	PrivateFirstClass
	Corporal
	SergeantThirdClass
	SergeantSecondClass
	SergeantFirstClass
	ChiefSergeant
	SecondLieutenant
	FirstLieutenant
	Captain
	SecondCommander
	FirstCommander
	HighCommander
	RearMarshal
	ViceMarshal
	Marshal
	GrandMarshal
	Champion
)

// Parse returns the primitive representation of the provided GC rank.
func Parse(input string) GCRank {
	switch input {
	case "Private Third Class":
		return PrivateThirdClass
	case "Private Second Class":
		return PrivateSecondClass
	case "Private First Class":
		return PrivateFirstClass
	case "Corporal":
		return Corporal
	case "Sergeant Third Class":
		return SergeantThirdClass
	case "Sergeant Second Class":
		return SergeantSecondClass
	case "Sergeant First Class":
		return SergeantFirstClass
	case "Chief Sergeant":
		return ChiefSergeant
	case "Second Lieutenant":
		return SecondLieutenant
	case "First Lieutenant":
		return FirstLieutenant
	case "Captain":
		return Captain
	case "Second Commander":
		return SecondCommander
	case "First Commander":
		return FirstCommander
	case "High Commander":
		return HighCommander
	case "Rear Marshal":
		return RearMarshal
	case "Vice Marshal":
		return ViceMarshal
	case "Marshal":
		return Marshal
	case "Grand Marshal":
		return GrandMarshal
	case "Champion":
		return Champion
	}
	return None
}
