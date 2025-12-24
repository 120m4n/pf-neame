package este

import (
	"fmt"
	"os"

	"github.com/120m4n/pf-neame/internal/utils"
	"github.com/spf13/cobra"
)

// EsteOptions holds the options for the este command
type EsteOptions struct {
	File    string
	Message string
	Row     int
	Column  int
	Output  string
}

// NewEsteCmd creates the este command
func NewEsteCmd() *cobra.Command {
	opts := &EsteOptions{}

	cmd := &cobra.Command{
		Use:   "este [filename]",
		Short: "Procesa archivos .exe, .dll, .pgi o .bpl para obtener su información de versión",
		Long:  `El comando este procesa archivos ejecutables y bibliotecas para extraer su información de versión.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.File = args[0]
			return runEste(opts)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&opts.Message, "message", "m", "", "Mensaje para procesar")
	cmd.Flags().IntVarP(&opts.Row, "row", "r", 0, "Número de fila")
	cmd.Flags().IntVarP(&opts.Column, "column", "c", 0, "Número de columna")
	cmd.Flags().StringVarP(&opts.Output, "output", "o", "", "Archivo de salida")

	return cmd
}

// runEste executes the este command logic
func runEste(opts *EsteOptions) error {
	return executeEste(opts)
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
	}

	// Validar que se proporcionó un archivo
	if opts.File == "" {
		return fmt.Errorf(humorousMessages["no_file"][0])
	}

	// Validar extensión del archivo
	if !utils.HasValidExtension(opts.File) {
		return fmt.Errorf(humorousMessages["invalid_extension"][0])
	}

	// Verificar si el archivo existe
	if _, err := os.Stat(opts.File); os.IsNotExist(err) {
		return fmt.Errorf(humorousMessages["file_not_found"][0])
	}

	// Obtener información de versión del archivo
	result := utils.GetFileVersionInfo(opts.File)
	
	// Mostrar información de versión
	fmt.Printf("=== Información de Versión para: %s ===\n\n", opts.File)
	
	if result.Info != nil {
		// Mostrar información disponible
		if result.Info.ProductVersion() != "" {
			fmt.Printf("Product Version: %s\n", result.Info.ProductVersion())
		}
		if result.Info.FileVersion() != "" {
			fmt.Printf("File Version: %s\n", result.Info.FileVersion())
		}
		if result.Info.CompanyName() != "" {
			fmt.Printf("Company Name: %s\n", result.Info.CompanyName())
		}
		if result.Info.FileDescription() != "" {
			fmt.Printf("File Description: %s\n", result.Info.FileDescription())
		}
		if result.Info.ProductName() != "" {
			fmt.Printf("Product Name: %s\n", result.Info.ProductName())
		}
		if result.Info.LegalCopyright() != "" {
			fmt.Printf("Legal Copyright: %s\n", result.Info.LegalCopyright())
		}
		if result.Info.OriginalFilename() != "" {
			fmt.Printf("Original Filename: %s\n", result.Info.OriginalFilename())
		}
		if result.Info.InternalName() != "" {
			fmt.Printf("Internal Name: %s\n", result.Info.InternalName())
		}
		if result.Info.Comments() != "" {
			fmt.Printf("Comments: %s\n", result.Info.Comments())
		}
	}
	
	// Si hay un error (por ejemplo, en Linux), mostrar advertencia con tono jocoso
	if result.Err != nil {
		fmt.Printf("\n⚠️  Nota: %v\n", result.Err)
		fmt.Println("¿Pero qué diablos haces usando Linux para esto? ¡Esto es para Windows, tonto!")
	}

	fmt.Println("\n¡Listo! Ahora deja de molestar y ponte a trabajar.")
	return nil
}
