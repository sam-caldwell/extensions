package extensions

import (
	"path/filepath"
	"strings"
)

// Get - Extract the file extension and return
func Get(name string) string {
	return strings.ToLower(filepath.Ext(name))
}
