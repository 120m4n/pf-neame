//go:build !windows
// +build !windows

package utils

import (
	"fmt"
)

// mockFileVersionInfo es una implementación mock para plataformas no-Windows
type mockFileVersionInfo struct {
	filename string
}

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

func (m *mockFileVersionInfo) ProductVersion() string    { return "N/A (Linux/Unix)" }
func (m *mockFileVersionInfo) FileVersion() string       { return "N/A (Linux/Unix)" }
func (m *mockFileVersionInfo) CompanyName() string       { return "N/A" }
func (m *mockFileVersionInfo) FileDescription() string   { return "File version info only available on Windows" }
func (m *mockFileVersionInfo) ProductName() string       { return "N/A" }
func (m *mockFileVersionInfo) LegalCopyright() string    { return "N/A" }
func (m *mockFileVersionInfo) OriginalFilename() string  { return m.filename }
func (m *mockFileVersionInfo) InternalName() string      { return "N/A" }
func (m *mockFileVersionInfo) Comments() string          { return "Run on Windows for full version info" }

// FileVersionResult contiene el resultado de obtener la información de versión
type FileVersionResult struct {
	Info FileVersionInfo
	Err  error
}

// GetFileVersionInfo obtiene la información de versión de un archivo
// En plataformas no-Windows, retorna un mock con información limitada
func GetFileVersionInfo(filename string) FileVersionResult {
	return FileVersionResult{
		Info: &mockFileVersionInfo{filename: filename},
		Err:  fmt.Errorf("file version extraction is only supported on Windows"),
	}
}
