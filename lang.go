package godestone

// Lang represents character language.
type Lang uint8

const (
	None Lang = 1 << iota
	JA
	EN
	DE
	FR
)
