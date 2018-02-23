package ry

import (
	"strings"
)

// Setting keeps track of
// - Site Name
// - Homepage
// - Layout
// - Theme
// - Site Privacy
// - Session Timeout
// - Layout
type Setting struct {
	Name        string
	Category    string
	Description string
	Value       string
}

// Settings is a slice of Setting{} and provides convenience methods for getting setting values
type Settings []Setting

func (s Settings) GetValue(settingName string) string {
	for _, v := range s {
		if strings.ToUpper(settingName) == strings.ToUpper(v.Value) {
			return v.Value
		}
	}
	return ""
}
