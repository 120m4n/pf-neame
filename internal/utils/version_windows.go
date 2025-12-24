//go:build windows

package utils

import (
	"errors"

	"github.com/bi-zone/go-fileversion"
)

// GetFileVersionInfo obtiene la información de versión de un archivo en Windows
func GetFileVersionInfo(filename string) FileVersionResult {
	info, err := fileversion.New(filename)

	if err == nil && len(info.Locales) == 0 {
		err = errors.New("upps!. parece que no hay información disponible")
	}

	return FileVersionResult{
		Info: info,
		Err:  err,
	}
}
