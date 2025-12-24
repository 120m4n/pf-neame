//go:build windows
// +build windows

package utils

import (
	"github.com/bi-zone/go-fileversion"
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

// GetFileVersionInfo obtiene la información de versión de un archivo en Windows
func GetFileVersionInfo(filename string) FileVersionResult {
	info, err := fileversion.New(filename)
	return FileVersionResult{
		Info: info,
		Err:  err,
	}
}
