package cmd

import (
	"fmt"
	"os"
	"strings"
	"uuid-converter/encoding"

	"github.com/spf13/cobra"
)

var (
	flagToLower bool
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode UUIDs returned by HEX MariaDB function",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing argument")
			os.Exit(1)
		}

		input := args[0]
		output, err := encoding.Decode(input)
		if err != nil {
			fmt.Printf("The decoding returned an error: %q\n", err)
			os.Exit(1)
		}

		if flagToLower {
			output = strings.ToLower(output)
		}

		fmt.Printf("Result: %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().BoolVar(&flagToLower, "lowercase", false, "Returns lowercase UUIDs")
}
