package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// PF26Data contiene los datos necesarios para editar el formato PF-26
type PF26Data struct {
	Product             string
	Client              string
	Version             string
	ProductDescription  string
	ClientDescription   string
	VersionDescription  string
}

// editCell es una función genérica no exportada que recibe el nombre del archivo,
// nombre de la hoja, fila y columna, y edita el valor de la celda especificada.
// Nota: Esta función abre y cierra el archivo para cada edición de celda, lo cual
// es ineficiente para ediciones múltiples. Se mantiene para casos de uso específicos
// donde se requiera editar una sola celda. Para ediciones múltiples, se recomienda
// usar EditPDF26 o implementar una función que mantenga el archivo abierto.
func editCell(fileName string, sheet string, row int, column string, value string) error {
	// Abrir el archivo Excel
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer func() {
		// Silently ignore close errors in defer as we've already saved the file
		_ = f.Close()
	}()

	// Construir la referencia de la celda (ej: "B11")
	cell := fmt.Sprintf("%s%d", column, row)

	// Establecer el valor de la celda
	if err := f.SetCellValue(sheet, cell, value); err != nil {
		return fmt.Errorf("error al establecer el valor de la celda %s: %w", cell, err)
	}

	// Guardar los cambios
	if err := f.Save(); err != nil {
		return fmt.Errorf("error al guardar el archivo: %w", err)
	}

	return nil
}

// EditPDF26 edita los valores del formato PF-26 en el archivo Excel especificado
// Recibe el nombre del archivo y una estructura con los datos a editar
// Si la hoja "FORMATO" no existe, la crea
func EditPDF26(fileName string, data PF26Data) error {
	// Abrir el archivo Excel
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer func() {
		// Silently ignore close errors in defer as we've already saved the file
		_ = f.Close()
	}()

	sheetName := "FORMATO"

	// Verificar si la hoja existe
	sheetIndex, err := f.GetSheetIndex(sheetName)
	if err != nil {
		return fmt.Errorf("error al obtener el índice de la hoja: %w", err)
	}

	// Si la hoja no existe (índice = -1), crearla
	if sheetIndex == -1 {
		_, err := f.NewSheet(sheetName)
		if err != nil {
			return fmt.Errorf("error al crear la hoja %s: %w", sheetName, err)
		}
	}

	// Editar las celdas con los valores proporcionados
	// Product -> Cell(B11)
	if err := f.SetCellValue(sheetName, "B11", data.Product); err != nil {
		return fmt.Errorf("error al establecer Product en B11: %w", err)
	}

	// Client -> Cell(B12)
	if err := f.SetCellValue(sheetName, "B12", data.Client); err != nil {
		return fmt.Errorf("error al establecer Client en B12: %w", err)
	}

	// Version -> Cell(B13)
	if err := f.SetCellValue(sheetName, "B13", data.Version); err != nil {
		return fmt.Errorf("error al establecer Version en B13: %w", err)
	}

	// Product_Description -> Cell(C11)
	if err := f.SetCellValue(sheetName, "C11", data.ProductDescription); err != nil {
		return fmt.Errorf("error al establecer ProductDescription en C11: %w", err)
	}

	// Client_Description -> Cell(C12)
	if err := f.SetCellValue(sheetName, "C12", data.ClientDescription); err != nil {
		return fmt.Errorf("error al establecer ClientDescription en C12: %w", err)
	}

	// Version_Description -> Cell(C13)
	if err := f.SetCellValue(sheetName, "C13", data.VersionDescription); err != nil {
		return fmt.Errorf("error al establecer VersionDescription en C13: %w", err)
	}

	// Guardar los cambios
	if err := f.Save(); err != nil {
		return fmt.Errorf("error al guardar el archivo: %w", err)
	}

	return nil
}
