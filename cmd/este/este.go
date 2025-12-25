package este

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/120m4n/pf-neame/internal/excel"
	"github.com/120m4n/pf-neame/internal/utils"
	"github.com/spf13/cobra"
)

// EsteOptions holds the options for the este command
type EsteOptions struct {
	File    string
	Message string
	Row     int
	Column  string
	Output  string
	Verbose bool
}

func getTemplatePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	filename := filepath.Join(filepath.Dir(ex), "./templates/pf-26.xlsx")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("el archivo de plantilla no existe en la ruta esperada: %s", filename)
	}

	return filename, nil
}

// NewEsteCmd creates the este command
func NewEsteCmd() *cobra.Command {
	opts := &EsteOptions{}

	cmd := &cobra.Command{
		Use:   "este [filename]",
		Short: "Le pasas un .exe, .dll, .pgi o .bpl y le gritas al QA !pf-neamee este ¡",
		Long:  `Le pasas un .exe, .dll, .pgi o .bpl y le gritas al QA !pf-neamee este. ¡Funciona solo en Windows!`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.File = args[0]
			return runEste(opts)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&opts.Message, "message", "m", "", "Mensaje para procesar")
	cmd.Flags().IntVarP(&opts.Row, "row", "r", 0, "Número de fila")
	cmd.Flags().StringVarP(&opts.Column, "column", "c", "", "Número de columna")
	cmd.Flags().StringVarP(&opts.Output, "output", "o", "", "Archivo de salida")
	// add verbose flag if needed
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "Habilitar salida detallada")

	return cmd
}

// runEste executes the este command logic
func runEste(opts *EsteOptions) error {
	return executeEste(opts)
}

// printIfNotEmpty imprime un campo de información solo si no está vacío
func printIfNotEmpty(label, value string) {
	if value != "" {
		fmt.Printf("%s: %s\n", label, value)
	}
}

// isValidColumn checks if the column string is a valid Excel column identifier (A-Z, AA-AZ)
func isValidColumn(col string) bool {
	if len(col) < 1 || len(col) > 2 {
		return false
	}
	for _, r := range col {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

// executeEste is the core business logic for the este command
func executeEste(opts *EsteOptions) error {
	// Lista de mensajes jocosos para diferentes situaciones
	humorousMessages := map[string][]string{
		"no_file": {
			"¿Eres tonto? ¡Tienes que especificar un archivo!",
			"Mira el puto letrero: pf-neame este <filename>",
			"¿Pero qué diablos haces? ¡Necesito un archivo!",
		},
		"invalid_extension": {
			"Esa extensión ni existe en mi lista. Solo acepto .exe, .dll, .pgi o .bpl",
			"¿Pero qué diablos haces con ese archivo? Necesito .exe, .dll, .pgi o .bpl",
			"Eres tonto o qué, eso no es un archivo válido. Lee las extensiones válidas: .exe, .dll, .pgi, .bpl",
		},
		"file_not_found": {
			"Esto no es Google, busca bien el archivo porque no existe",
			"¿Eres tonto? Ese archivo no existe en tu sistema",
			"Mira el puto letrero: el archivo no se encuentra",
		},
		"version_error": {
			"Esa función ni existe o el archivo está corrupto",
			"¿Pero qué diablos haces? No puedo leer la versión de ese archivo",
			"Eres tonto, el archivo no tiene información de versión o está roto",
		},
		"unknown_error": {
			"Algo raro pasó, no sé qué diablos fue",
			"¿Eres tonto? No entiendo qué pasó aquí",
			"Mira el puto letrero: ocurrió un error desconocido",
		},
	}

	if opts.Row < 0 {
		return fmt.Errorf("El número de fila y columna no puede ser negativo")
	}

	// Editar el archivo Excel con los datos obtenidos
	templatePath, err := getTemplatePath()
	if err != nil {
		return fmt.Errorf("parece que lo que quieres es hacer magia: %w", err)
	}

	// Inicializar generador de números aleatorios
	// para seleccionar mensajes jocosos
	indexRandom := rand.Intn(3)

	opts.Column = strings.ToUpper(opts.Column)

	if opts.Row > 0 && opts.Column != "" && opts.Message != "" && isValidColumn(opts.Column) {
		err := excel.EditCell(templatePath, "FORMATO", opts.Row, opts.Column, opts.Message)
		if err != nil {
			return fmt.Errorf(humorousMessages["no_file"][indexRandom])
		}
	}

	// Validar que se proporcionó un archivo
	if opts.File == "" {
		return fmt.Errorf(humorousMessages["no_file"][indexRandom])
	}

	// Validar extensión del archivo
	if !utils.HasValidExtension(opts.File) {
		return fmt.Errorf(humorousMessages["invalid_extension"][indexRandom])
	}

	// Verificar si el archivo existe
	if _, err := os.Stat(opts.File); os.IsNotExist(err) {
		return fmt.Errorf(humorousMessages["unknown_error"][indexRandom])
	}

	if opts.Column != "" && !isValidColumn(opts.Column) {
		return fmt.Errorf("Columna inválida: %s", opts.Column)
	}

	// Obtener información de versión del archivo
	result := utils.GetFileVersionInfo(opts.File)

	// Si hay un error (por ejemplo, en Linux), mostrar advertencia con tono jocoso
	if result.Err != nil {
		fmt.Printf("\n⚠️  Nota: %v\n", result.Err)
		return fmt.Errorf(humorousMessages["version_error"][indexRandom])
	}

	// Mostrar información de versión
	fmt.Printf("=== Información de Versión para: %s ===\n\n", opts.File)

	if result.Info != nil && opts.Verbose {
		// Mostrar información disponible
		printIfNotEmpty("Product Version", result.Info.ProductVersion())
		printIfNotEmpty("File Version", result.Info.FileVersion())
		printIfNotEmpty("Company Name", result.Info.CompanyName())
		printIfNotEmpty("File Description", result.Info.FileDescription())
		printIfNotEmpty("Product Name", result.Info.ProductName())
		printIfNotEmpty("Legal Copyright", result.Info.LegalCopyright())
		printIfNotEmpty("Original Filename", result.Info.OriginalFilename())
		printIfNotEmpty("Internal Name", result.Info.InternalName())
		printIfNotEmpty("Comments", result.Info.Comments())
	}

	// Chequear que comments no esté vacío antes de proceder
	comentarios := strings.TrimSpace(result.Info.Comments())
	if comentarios == "" {
		comentarios = "WTF! Sin comentarios"
	}

	data := excel.PF26Data{
		Product:            "1",
		Client:             "",
		Version:            result.Info.FileVersion(),
		ProductDescription: result.Info.ProductName(),
		ClientDescription:  "ESSA",
		VersionDescription: comentarios,
	}

	err = excel.EditPDF26(templatePath, data)
	if err != nil {
		return fmt.Errorf("error al editar el archivo Excel: %w", err)
	}

	fmt.Println("\n¡Listo! Ahora deja de molestar y ponte a trabajar.")
	return nil
}
