package godestone

// SearchOrder represents the search result ordering of a Lodestone search.
type SearchOrder uint8

const (
	OrderNameAToZ        SearchOrder = 1
	OrderNameZToA        SearchOrder = 2
	OrderWorldAtoZ       SearchOrder = 3
	OrderWorldZtoA       SearchOrder = 4
	OrderLevelDescending SearchOrder = 5
	OrderLevelAscending  SearchOrder = 6
)
