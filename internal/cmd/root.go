package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "konstraint",
	Short:   "",
	Version: "0.1.0",
}

func Execute() {
	rootCmd.AddCommand(TemplateCli())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
