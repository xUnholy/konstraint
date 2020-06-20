package cmd

import "github.com/spf13/cobra"

func TemplateCli() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "template",
		Short: "",
		Run:   templateCmd,
	}
	return startCmd
}

func templateCmd(cmd *cobra.Command, args []string) {

}
