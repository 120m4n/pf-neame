package main

import (
	"fmt"
	"os"

	"github.com/120m4n/pf-neame/cmd/este"
	"github.com/120m4n/pf-neame/internal/version"
	"github.com/spf13/cobra"
)

// newRootCmd creates the root command for pf-neame CLI
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "pf-neame",
		Short:   "Utilidad CLI para que el QA se pf-nemee archivos .dll o .exe",
		Version: version.Version,
	}

	// Add subcommands
	rootCmd.AddCommand(este.NewEsteCmd())

	return rootCmd
}

func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
