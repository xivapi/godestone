package godestone

// Lang represents character language.
type Lang uint8

const (
	JA Lang = 1 << iota
	EN
	DE
	FR
)
