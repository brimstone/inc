package main

import (
	"fmt"
	"os"

	clientCmd "github.com/brimstone/inc/cmd/inc-client/cmd"
	serverCmd "github.com/brimstone/inc/cmd/incd/cmd"
	"github.com/brimstone/inc/pkg/cmd"
	"github.com/brimstone/inc/pkg/version"
	"github.com/spf13/cobra"
)

func main() {

	version.Binary = "inc"
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   version.Binary,
		Short: "All in one inc binary",
		Long:  `This is both a client and server for Inc.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}
	var cfgFile string
	// Initalization
	cobra.OnInitialize(cmd.InitConfig(&cfgFile))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.inc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmd.AddAll(rootCmd)
	clientCmd.AddAll(rootCmd)
	serverCmd.AddAll(rootCmd)
	// Former Execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
