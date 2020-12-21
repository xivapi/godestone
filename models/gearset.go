package models

// GearItem represents information about a single gear item on a character.
type GearItem struct {
	Creator string
	Dye     uint32
	ID      uint32
	Materia []uint32
	Mirage  uint32
}

// GearItemBuild represents a full gearset on a character. All gear items can be nil.
type GearItemBuild struct {
	Body        *GearItem
	Bracelets   *GearItem
	Earrings    *GearItem
	Feet        *GearItem
	Hands       *GearItem
	Head        *GearItem
	Legs        *GearItem
	MainHand    *GearItem
	Necklace    *GearItem
	OffHand     *GearItem
	Ring1       *GearItem
	Ring2       *GearItem
	SoulCrystal *GearItem
	Waist       *GearItem
}

// GearSet represents the current gear information of a character.
type GearSet struct {
	Attributes map[uint8]uint32
	ClassID    uint8
	Gear       *GearItemBuild
	GearKey    string
	JobID      uint8
	Level      uint8
}
