package search

import "strings"

// ActiveMemberRange represents the active member range filter of a search.
type ActiveMemberRange string

// Active member range for searches.
const (
	OneToTen         ActiveMemberRange = "1-10"
	ElevenToThirty   ActiveMemberRange = "11-30"
	ThirtyOneToFifty ActiveMemberRange = "31-50"
	FiftyOnePlus     ActiveMemberRange = "51-"
)

func parseWorldDC(world string, dc string) string {
	worldDC := dc
	if len(world) != 0 {
		worldDC = world
	} else {
		// DCs have the _dc_ prefix attached to them
		if len(worldDC) != 0 && !strings.HasPrefix(worldDC, "_dc_") {
			worldDC = "_dc_" + worldDC
		}
	}
	return worldDC
}
