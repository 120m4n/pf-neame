# Excel Package - Ejemplo de Uso

Este archivo contiene un ejemplo simple de cómo usar el paquete `excel` para editar documentos PF-26.

## Ejemplo Básico

```go
package main

import (
    "fmt"
    "log"

    "github.com/120m4n/pf-neame/internal/excel"
)

func main() {
    // Preparar los datos para el documento PF-26
    data := excel.PF26Data{
        Product:             "MiAplicacion.exe",
        Client:              "Empresa ABC S.A.",
        Version:             "1.0.0",
        ProductDescription:  "Sistema de gestión empresarial integrado",
        ClientDescription:   "Empresa dedicada a servicios financieros",
        VersionDescription:  "Primera versión estable del sistema",
    }

    // Editar el documento PF-26
    // Si el archivo no existe, asegúrate de crearlo primero con Excel
    err := excel.EditPDF26("documento_pf26.xlsx", data)
    if err != nil {
        log.Fatalf("Error al editar el documento: %v", err)
    }

    fmt.Println("✓ Documento PF-26 actualizado exitosamente")
    fmt.Println("  - Product: B11")
    fmt.Println("  - Client: B12")
    fmt.Println("  - Version: B13")
    fmt.Println("  - Product Description: C11")
    fmt.Println("  - Client Description: C12")
    fmt.Println("  - Version Description: C13")
}
```

## Integración con el comando "este"

Puedes integrar esta funcionalidad con el comando `este` para automatizar el proceso:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/120m4n/pf-neame/internal/excel"
    "github.com/120m4n/pf-neame/internal/utils"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Uso: programa <archivo.exe|dll|pgi|bpl> <documento_pf26.xlsx>")
    }

    binFile := os.Args[1]
    excelFile := "PF-26.xlsx"
    if len(os.Args) >= 3 {
        excelFile = os.Args[2]
    }

    // Obtener información del archivo binario
    result := utils.GetFileVersionInfo(binFile)
    if result.Err != nil {
        log.Fatalf("Error al obtener información de versión: %v", result.Err)
    }

    // Preparar datos desde la información del archivo
    data := excel.PF26Data{
        Product:             result.Info.ProductName(),
        Client:              result.Info.CompanyName(),
        Version:             result.Info.ProductVersion(),
        ProductDescription:  result.Info.FileDescription(),
        ClientDescription:   result.Info.Comments(),
        VersionDescription:  fmt.Sprintf("Versión de archivo: %s", result.Info.FileVersion()),
    }

    // Editar el documento PF-26
    err := excel.EditPDF26(excelFile, data)
    if err != nil {
        log.Fatalf("Error al editar el documento PF-26: %v", err)
    }

    fmt.Printf("✓ Documento PF-26 '%s' generado automáticamente desde '%s'\n", excelFile, binFile)
}
```

## Notas Importantes

1. El archivo Excel debe existir antes de usar `EditPDF26`
2. Si la hoja "FORMATO" no existe, se creará automáticamente
3. Los valores existentes en las celdas serán sobrescritos
4. Todos los errores se retornan con contexto descriptivo

## Crear un Archivo Excel Base

Si necesitas crear un archivo Excel base, puedes usar este código:

```go
package main

import (
    "log"

    "github.com/xuri/excelize/v2"
)

func main() {
    // Crear un nuevo archivo Excel
    f := excelize.NewFile()
    
    // Crear la hoja FORMATO
    _, err := f.NewSheet("FORMATO")
    if err != nil {
        log.Fatalf("Error al crear la hoja: %v", err)
    }
    
    // Establecer encabezados (opcional)
    f.SetCellValue("FORMATO", "A11", "Producto:")
    f.SetCellValue("FORMATO", "A12", "Cliente:")
    f.SetCellValue("FORMATO", "A13", "Versión:")
    
    // Guardar el archivo
    if err := f.SaveAs("documento_pf26_base.xlsx"); err != nil {
        log.Fatalf("Error al guardar: %v", err)
    }
    
    f.Close()
    println("✓ Archivo base creado exitosamente")
}
```

Para más información, consulta el archivo CASOS_DE_USO.md en el directorio internal/excel/.
