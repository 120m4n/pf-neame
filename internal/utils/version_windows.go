//go:build windows

package utils

import (
	"github.com/bi-zone/go-fileversion"
)

// GetFileVersionInfo obtiene la información de versión de un archivo en Windows
func GetFileVersionInfo(filename string) FileVersionResult {
	info, err := fileversion.New(filename)
	return FileVersionResult{
		Info: info,
		Err:  err,
	}
}
