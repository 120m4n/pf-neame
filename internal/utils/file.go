package utils

import (
	"path/filepath"
	"strings"
)

// HasValidExtension verifica si el archivo tiene una extensión válida
// Extensiones válidas: .exe, .dll, .pgi, .bpl
func HasValidExtension(filename string) bool {
	validExtensions := []string{".exe", ".dll", ".pgi", ".bpl"}
	ext := strings.ToLower(filepath.Ext(filename))
	
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	
	return false
}
