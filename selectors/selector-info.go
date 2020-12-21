package selectors

// SelectorInfo represents a variety of information about a CSS selector.
type SelectorInfo struct {
	Selector string `json:"selector"`
	Regex    string `json:"regex,omitempty"`
}
