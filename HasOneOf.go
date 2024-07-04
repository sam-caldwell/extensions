package extensions

import "strings"

func HasOneOf(name string, expectedExtensions ...string) bool {
	for _, ext := range expectedExtensions {
		if Get(name) == strings.ToLower(ext) {
			return true
		}
	}
	return false
}
