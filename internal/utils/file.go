package utils

import (
	"path/filepath"
	"strings"
)

// FileVersionInfo representa la información de versión de un archivo
type FileVersionInfo interface {
	ProductVersion() string
	FileVersion() string
	CompanyName() string
	FileDescription() string
	ProductName() string
	LegalCopyright() string
	OriginalFilename() string
	InternalName() string
	Comments() string
}

// FileVersionResult contiene el resultado de obtener la información de versión
type FileVersionResult struct {
	Info FileVersionInfo
	Err  error
}

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
