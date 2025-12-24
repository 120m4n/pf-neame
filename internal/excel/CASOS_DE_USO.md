# Casos de Uso - Paquete Excel

Este documento describe los casos de uso para el paquete `excel` que proporciona funcionalidades para editar archivos Excel utilizando la librería `excelize/v2`.

## Descripción General

El paquete `excel` proporciona funciones para editar archivos Excel, específicamente diseñado para automatizar el llenado de documentos PF-26 utilizados en procesos de QA.

## Funciones Disponibles

### 1. `EditPDF26(fileName string, data PF26Data) error`

Función exportada que edita los valores del formato PF-26 en un archivo Excel.

#### Parámetros

- `fileName`: Ruta del archivo Excel a editar (string)
- `data`: Estructura PF26Data con los datos a insertar

#### Estructura PF26Data

```go
type PF26Data struct {
    Product             string  // Nombre del producto
    Client              string  // Nombre del cliente
    Version             string  // Versión del producto
    ProductDescription  string  // Descripción del producto
    ClientDescription   string  // Descripción del cliente
    VersionDescription  string  // Descripción de la versión
}
```

#### Mapeo de Celdas

La función mapea los datos a las siguientes celdas en la hoja "FORMATO":

| Campo | Celda | Descripción |
|-------|-------|-------------|
| Product | B11 | Nombre del producto |
| Client | B12 | Nombre del cliente |
| Version | B13 | Versión del producto |
| ProductDescription | C11 | Descripción del producto |
| ClientDescription | C12 | Descripción del cliente |
| VersionDescription | C13 | Descripción de la versión |

#### Comportamiento

- Si la hoja "FORMATO" no existe en el archivo, la función la crea automáticamente
- Si el archivo no existe, retorna un error
- Todos los errores son retornados con contexto descriptivo

## Casos de Uso

### Caso de Uso 1: Llenar un documento PF-26 nuevo

**Descripción**: Como usuario del sistema de QA, quiero llenar automáticamente un documento PF-26 con la información de un nuevo producto.

**Precondiciones**:
- Existe un archivo Excel con formato PF-26 (puede no tener la hoja FORMATO)

**Flujo Principal**:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/120m4n/pf-neame/internal/excel"
)

func main() {
    data := excel.PF26Data{
        Product:             "MiAplicacion.exe",
        Client:              "Cliente ABC",
        Version:             "1.0.0",
        ProductDescription:  "Aplicación de gestión empresarial",
        ClientDescription:   "Empresa de servicios financieros",
        VersionDescription:  "Primera versión estable del producto",
    }
    
    err := excel.EditPDF26("documento_pf26.xlsx", data)
    if err != nil {
        log.Fatalf("Error al editar el documento: %v", err)
    }
    
    fmt.Println("Documento PF-26 actualizado exitosamente")
}
```

**Postcondiciones**:
- El archivo Excel contiene la hoja "FORMATO" con los datos especificados
- Todas las celdas (B11-B13 y C11-C13) contienen los valores proporcionados

---

### Caso de Uso 2: Actualizar información existente en un PF-26

**Descripción**: Como usuario del sistema de QA, quiero actualizar la información de versión en un documento PF-26 existente.

**Precondiciones**:
- Existe un archivo Excel con la hoja "FORMATO" que ya contiene datos

**Flujo Principal**:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/120m4n/pf-neame/internal/excel"
)

func main() {
    data := excel.PF26Data{
        Product:             "MiAplicacion.exe",
        Client:              "Cliente ABC",
        Version:             "2.0.0",
        ProductDescription:  "Aplicación de gestión empresarial",
        ClientDescription:   "Empresa de servicios financieros",
        VersionDescription:  "Segunda versión con nuevas funcionalidades",
    }
    
    err := excel.EditPDF26("documento_pf26_existente.xlsx", data)
    if err != nil {
        log.Fatalf("Error al actualizar el documento: %v", err)
    }
    
    fmt.Println("Documento PF-26 actualizado con nueva versión")
}
```

**Postcondiciones**:
- Los valores en las celdas son reemplazados por los nuevos valores proporcionados
- El resto del contenido del archivo permanece intacto

---

### Caso de Uso 3: Crear hoja FORMATO en documento existente

**Descripción**: Como usuario del sistema de QA, quiero agregar la hoja FORMATO a un documento Excel existente que no la tiene.

**Precondiciones**:
- Existe un archivo Excel sin la hoja "FORMATO"

**Flujo Principal**:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/120m4n/pf-neame/internal/excel"
)

func main() {
    data := excel.PF26Data{
        Product:             "NuevoProducto.dll",
        Client:              "Cliente XYZ",
        Version:             "1.5.3",
        ProductDescription:  "Librería de utilidades compartidas",
        ClientDescription:   "Empresa de desarrollo de software",
        VersionDescription:  "Versión con correcciones de seguridad",
    }
    
    err := excel.EditPDF26("documento_sin_formato.xlsx", data)
    if err != nil {
        log.Fatalf("Error al editar el documento: %v", err)
    }
    
    fmt.Println("Hoja FORMATO creada y datos agregados exitosamente")
}
```

**Postcondiciones**:
- Se crea la hoja "FORMATO" en el archivo
- La hoja contiene los datos especificados en las celdas correspondientes

---

### Caso de Uso 4: Integración con el comando "este"

**Descripción**: Como desarrollador del sistema, quiero integrar la funcionalidad de edición de Excel con el comando "este" para automatizar completamente el proceso de documentación.

**Precondiciones**:
- Se ha ejecutado el comando "este" sobre un archivo .exe/.dll/.pgi/.bpl
- Se ha obtenido la información de versión del archivo

**Flujo Principal**:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/120m4n/pf-neame/internal/excel"
    "github.com/120m4n/pf-neame/internal/utils"
)

func main() {
    // Obtener información del archivo
    fileName := "MiAplicacion.exe"
    result := utils.GetFileVersionInfo(fileName)
    
    if result.Err != nil {
        log.Fatalf("Error al obtener información de versión: %v", result.Err)
    }
    
    // Preparar datos para el documento PF-26
    data := excel.PF26Data{
        Product:             result.Info.ProductName(),
        Client:              result.Info.CompanyName(),
        Version:             result.Info.ProductVersion(),
        ProductDescription:  result.Info.FileDescription(),
        ClientDescription:   result.Info.Comments(),
        VersionDescription:  fmt.Sprintf("Versión de archivo: %s", result.Info.FileVersion()),
    }
    
    // Editar el documento PF-26
    err := excel.EditPDF26("PF-26.xlsx", data)
    if err != nil {
        log.Fatalf("Error al editar el documento PF-26: %v", err)
    }
    
    fmt.Println("Documento PF-26 generado automáticamente desde el archivo")
}
```

**Postcondiciones**:
- El documento PF-26 se llena automáticamente con la información extraída del archivo binario
- El proceso de documentación se automatiza completamente

---

### Caso de Uso 5: Procesamiento por lotes

**Descripción**: Como usuario del sistema de QA, quiero procesar múltiples documentos PF-26 en un solo proceso.

**Precondiciones**:
- Existen múltiples archivos Excel que necesitan ser actualizados

**Flujo Principal**:
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/120m4n/pf-neame/internal/excel"
)

func main() {
    documentos := []struct {
        archivo string
        datos   excel.PF26Data
    }{
        {
            archivo: "PF-26_Producto1.xlsx",
            datos: excel.PF26Data{
                Product:             "Producto1.exe",
                Client:              "Cliente A",
                Version:             "1.0.0",
                ProductDescription:  "Descripción del Producto 1",
                ClientDescription:   "Descripción del Cliente A",
                VersionDescription:  "Primera versión",
            },
        },
        {
            archivo: "PF-26_Producto2.xlsx",
            datos: excel.PF26Data{
                Product:             "Producto2.dll",
                Client:              "Cliente B",
                Version:             "2.0.0",
                ProductDescription:  "Descripción del Producto 2",
                ClientDescription:   "Descripción del Cliente B",
                VersionDescription:  "Segunda versión",
            },
        },
    }
    
    for _, doc := range documentos {
        err := excel.EditPDF26(doc.archivo, doc.datos)
        if err != nil {
            log.Printf("Error al editar %s: %v", doc.archivo, err)
            continue
        }
        fmt.Printf("✓ Documento %s actualizado exitosamente\n", doc.archivo)
    }
}
```

**Postcondiciones**:
- Todos los documentos son procesados
- Se reportan errores individuales sin interrumpir el procesamiento del resto

---

## Manejo de Errores

La función `EditPDF26` puede retornar los siguientes tipos de errores:

1. **Error al abrir el archivo**: El archivo no existe o no se tienen permisos de lectura
   ```
   error al abrir el archivo: <detalle>
   ```

2. **Error al obtener el índice de la hoja**: Problema interno al verificar si la hoja existe
   ```
   error al obtener el índice de la hoja: <detalle>
   ```

3. **Error al crear la hoja**: No se pudo crear la hoja "FORMATO"
   ```
   error al crear la hoja FORMATO: <detalle>
   ```

4. **Error al establecer valor de celda**: No se pudo escribir en una celda específica
   ```
   error al establecer <campo> en <celda>: <detalle>
   ```

5. **Error al guardar el archivo**: No se pudieron guardar los cambios
   ```
   error al guardar el archivo: <detalle>
   ```

## Mejores Prácticas

1. **Siempre manejar errores**: Verificar el error retornado por `EditPDF26`
2. **Validar datos de entrada**: Asegurarse de que los campos de `PF26Data` no estén vacíos si son obligatorios
3. **Backup de archivos**: Considerar crear una copia de respaldo antes de editar archivos importantes
4. **Permisos de archivo**: Asegurarse de tener permisos de escritura en el archivo Excel
5. **Rutas absolutas**: Usar rutas absolutas para evitar problemas con el directorio de trabajo

## Limitaciones

1. La función solo edita la hoja "FORMATO"
2. Las celdas de destino están fijas (B11-B13, C11-C13)
3. No se valida el formato del archivo Excel de entrada
4. No se crean estructuras adicionales en la hoja (solo se escriben valores)

## Notas Técnicas

- La función utiliza la librería `github.com/xuri/excelize/v2` para manipular archivos Excel
- Los archivos Excel deben estar en formato .xlsx
- La función `editCell` es privada y no puede ser utilizada directamente desde fuera del paquete
- El cierre del archivo se maneja mediante `defer` para garantizar que se cierre incluso si hay errores
