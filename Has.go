package extensions

import "strings"

// Has - Return boolean true if the given name has the given file extension.
func Has(name string, expectedExtension string) bool {
	return Get(name) == strings.ToLower(expectedExtension)
}
