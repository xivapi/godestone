package models

// Meta is meta.
type Meta struct {
	Version          string `json:"version"`
	UserAgentDesktop string `json:"userAgentDesktop"`
	UserAgentMobile  string `json:"userAgentMobile"`
}
