package cmd

import "github.com/spf13/cobra"

func AddAll(rootCmd *cobra.Command) {
	rootCmd.AddCommand(SubmitCmd())
}
