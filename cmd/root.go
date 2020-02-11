package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "test [command]",
	Short: "This is a test",
	Long:  "This is a test",
}

// Execute root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("command", "default-command", "Command to run")
	viper.BindPFlag("command", rootCmd.PersistentFlags().Lookup("command"))
}

func initConfig() {
	viper.SetEnvPrefix("test")
	viper.AutomaticEnv()
}
