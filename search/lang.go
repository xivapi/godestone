package search

// Lang represents character language.
type Lang uint8

// Language
const (
	NoneLang Lang = 1 << iota
	JA
	EN
	DE
	FR
)
