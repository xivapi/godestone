package selectors

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// SelectorInfo represents a variety of information about a CSS selector.
type SelectorInfo struct {
	Selector  string `json:"selector"`
	Regex     string `json:"regex,omitempty"`
	Attribute string `json:"attribute,omitempty"`

	loadedRegex *regexp.Regexp
}

func transformRegexForGolang(regex string) string {
	return strings.Replace(regex, "\\s+", "[\\s\\x{00A0}]+", -1)
}

func (si *SelectorInfo) runRegexIfExists(text string) []string {
	if si.Regex == "" {
		return nil
	}

	if si.loadedRegex == nil {
		transformedRegex := transformRegexForGolang(si.Regex)
		si.loadedRegex = regexp.MustCompile(transformedRegex)
	}

	matches := si.loadedRegex.FindStringSubmatch(text)
	return matches
}

// Parse returns values from an HTML element.
func (si *SelectorInfo) Parse(e *colly.HTMLElement) []string {
	text := e.Text
	if si.Attribute != "" {
		text = e.Attr(si.Attribute)
	}

	matches := si.runRegexIfExists(text)
	if matches == nil {
		return []string{text}
	}
	return matches[1:]
}

// ParseInnerHTML returns values from an HTML element.
func (si *SelectorInfo) ParseInnerHTML(e *colly.HTMLElement) []string {
	text, err := e.DOM.Html()
	if err != nil { // I'm doing something wrong if this is hit
		panic(err)
	}

	if si.Attribute != "" {
		text = e.Attr(si.Attribute)
	}

	matches := si.runRegexIfExists(text)
	if matches == nil {
		return []string{text}
	}
	return matches[1:]
}

// ParseSelection returns values from a selection.
func (si *SelectorInfo) ParseSelection(e *goquery.Selection) []string {
	text := e.Text()
	if si.Attribute != "" {
		attrText, exists := e.Attr(si.Attribute)
		if exists {
			text = attrText
		}
	}

	matches := si.runRegexIfExists(text)
	if matches == nil {
		return []string{text}
	}
	return matches[1:]
}

// ParseThroughChildren returns values from somewhere in an HTML element tree.
func (si *SelectorInfo) ParseThroughChildren(e *colly.HTMLElement) []string {
	el := e.DOM.Find(si.Selector)
	texts := si.ParseSelection(el)
	return texts
}
