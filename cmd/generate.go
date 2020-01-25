package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate generates output",
	Long: `Generates output...

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Call the busn logic
	},

	Args: func(cmd *cobra.Command, args []string) error {
		// Check any arguments
		return nil
	},
}

// Flag is the server's port
// var Flag int

func init() {
	RootCmd.AddCommand(generateCmd)
	// flags := generateCmd.PersistentFlags()

	// flags.StringVarP(&Fixdir, "fixdir", "", ".", "the directory to operate in")
	// flags.IntVarP(&Port, "port", "p", 8088, "The port to listen on for requests")
	// flags.StringSliceVarP(&Modulepath, "modulepath", "", []string{}, "one or more modulepaths to search instead of <fixdir>/modules")
	// flags.BoolVar(&Unsafe, "unsafe", false, "If unsafe http 1.1 without TLS should be used. Do not use for production!")
	// flags.StringVarP(&CaFile, "ca", "", "", "The location of the CA")
}
