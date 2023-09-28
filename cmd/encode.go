package cmd

import (
	"fmt"
	"os"
	"uuid-converter/encoding"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode UUIDs to be ready to be passed to UNHEX MariaDB function",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing argument")
			os.Exit(1)
		}

		input := args[0]
		output, err := encoding.Encode(input)
		if err != nil {
			fmt.Printf("The encoding returned an error: %q\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
