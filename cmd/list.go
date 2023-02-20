package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current configuration values.",
	Long:  `List the current configuration values.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("Configurations")
		for k, v := range viper.AllSettings() {
			fmt.Printf("%s: %s\n", color.GreenString(strings.Title(k)), v)
		}
	},
}

func init() {
	configCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
