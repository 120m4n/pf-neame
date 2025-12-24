package excel

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

// TestEditPDF26_NewSheet prueba la creación de una nueva hoja FORMATO
func TestEditPDF26_NewSheet(t *testing.T) {
	// Crear un directorio temporal
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test_pf26_new.xlsx"

	// Crear un archivo Excel básico
	f := excelize.NewFile()
	if err := f.SaveAs(tmpFile); err != nil {
		t.Fatalf("Error al crear archivo temporal: %v", err)
	}
	f.Close()

	// Datos de prueba
	data := PF26Data{
		Product:             "TestApp.exe",
		Client:              "Test Client",
		Version:             "1.0.0",
		ProductDescription:  "Test application description",
		ClientDescription:   "Test client description",
		VersionDescription:  "Initial test version",
	}

	// Ejecutar EditPDF26
	err := EditPDF26(tmpFile, data)
	if err != nil {
		t.Fatalf("Error al ejecutar EditPDF26: %v", err)
	}

	// Verificar que el archivo fue modificado correctamente
	f, err = excelize.OpenFile(tmpFile)
	if err != nil {
		t.Fatalf("Error al abrir archivo modificado: %v", err)
	}
	defer f.Close()

	// Verificar que la hoja FORMATO existe
	sheetIndex, _ := f.GetSheetIndex("FORMATO")
	if sheetIndex == -1 {
		t.Error("La hoja FORMATO no fue creada")
	}

	// Verificar los valores de las celdas
	tests := []struct {
		cell     string
		expected string
	}{
		{"B11", data.Product},
		{"B12", data.Client},
		{"B13", data.Version},
		{"C11", data.ProductDescription},
		{"C12", data.ClientDescription},
		{"C13", data.VersionDescription},
	}

	for _, tt := range tests {
		value, err := f.GetCellValue("FORMATO", tt.cell)
		if err != nil {
			t.Errorf("Error al obtener valor de celda %s: %v", tt.cell, err)
			continue
		}
		if value != tt.expected {
			t.Errorf("Celda %s: esperado %q, obtenido %q", tt.cell, tt.expected, value)
		}
	}
}

// TestEditPDF26_ExistingSheet prueba la edición de una hoja FORMATO existente
func TestEditPDF26_ExistingSheet(t *testing.T) {
	// Crear un directorio temporal
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test_pf26_existing.xlsx"

	// Crear un archivo Excel con la hoja FORMATO
	f := excelize.NewFile()
	_, err := f.NewSheet("FORMATO")
	if err != nil {
		t.Fatalf("Error al crear hoja FORMATO: %v", err)
	}
	if err := f.SaveAs(tmpFile); err != nil {
		t.Fatalf("Error al guardar archivo temporal: %v", err)
	}
	f.Close()

	// Datos de prueba
	data := PF26Data{
		Product:             "UpdatedApp.dll",
		Client:              "Updated Client",
		Version:             "2.0.0",
		ProductDescription:  "Updated application description",
		ClientDescription:   "Updated client description",
		VersionDescription:  "Second version",
	}

	// Ejecutar EditPDF26
	err = EditPDF26(tmpFile, data)
	if err != nil {
		t.Fatalf("Error al ejecutar EditPDF26: %v", err)
	}

	// Verificar que el archivo fue modificado correctamente
	f, err = excelize.OpenFile(tmpFile)
	if err != nil {
		t.Fatalf("Error al abrir archivo modificado: %v", err)
	}
	defer f.Close()

	// Verificar los valores de las celdas
	tests := []struct {
		cell     string
		expected string
	}{
		{"B11", data.Product},
		{"B12", data.Client},
		{"B13", data.Version},
		{"C11", data.ProductDescription},
		{"C12", data.ClientDescription},
		{"C13", data.VersionDescription},
	}

	for _, tt := range tests {
		value, err := f.GetCellValue("FORMATO", tt.cell)
		if err != nil {
			t.Errorf("Error al obtener valor de celda %s: %v", tt.cell, err)
			continue
		}
		if value != tt.expected {
			t.Errorf("Celda %s: esperado %q, obtenido %q", tt.cell, tt.expected, value)
		}
	}
}

// TestEditPDF26_FileNotFound prueba el manejo de error cuando el archivo no existe
func TestEditPDF26_FileNotFound(t *testing.T) {
	// Usar un directorio temporal para un archivo inexistente
	tmpDir := t.TempDir()
	nonExistentFile := tmpDir + "/nonexistent_file.xlsx"
	
	data := PF26Data{
		Product: "Test",
		Client:  "Test",
		Version: "1.0",
	}

	err := EditPDF26(nonExistentFile, data)
	if err == nil {
		t.Error("Se esperaba un error para archivo inexistente, pero no se obtuvo ninguno")
	}
}

// TestEditCell prueba la función privada editCell indirectamente
// a través del comportamiento de EditPDF26
func TestEditCell(t *testing.T) {
	// Esta función privada se prueba indirectamente a través de EditPDF26
	// que la utiliza internamente. Los tests de EditPDF26 cubren su funcionalidad.
	t.Skip("editCell es una función privada, se prueba a través de EditPDF26")
}
