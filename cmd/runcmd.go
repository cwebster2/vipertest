package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run  [command]",
	Short: "This is a test",
	Long:  "This is a test",
	Run:   runHandler,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runHandler(cmd *cobra.Command, args []string) {
	teststr := viper.GetString("command")
	fmt.Println("The command to run is: ", teststr)
}
