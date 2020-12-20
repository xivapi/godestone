package models

// ClassJob represents class and job information.
type ClassJob struct {
	ClassID       uint8
	ExpLevel      uint32
	ExpLevelMax   uint32
	ExpLevelTogo  uint32
	IsSpecialized bool
	JobID         uint8
	Level         uint8
	Name          string
	UnlockedState UnlockedState
}

// ClassJobBozja represents character progression data in the Bozjan Southern Front.
type ClassJobBozja struct {
	Level  uint8
	Mettle uint32
	Name   string
}

// ClassJobEureka represents character progression data in Eureka.
type ClassJobEureka struct {
	ExpLevel     uint32
	ExpLevelMax  uint32
	ExpLevelTogo uint32
	Level        uint8
	Name         string
}

// UnlockedState represents the unlock state of a ClassJob
type UnlockedState struct {
	ID   uint8
	Name string
}
