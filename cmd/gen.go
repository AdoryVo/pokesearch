package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/adoryvo/pokesearch/util"
)

var genCmd = &cobra.Command{
	Use:   "gen [gen #]",
	Short: "Set your default gen for searching learnsets & other information.",
	Long:  `Set your default gen for searching learnsets & other information.`,
	Example: `
* pokesearch config gen 4 -> change default gen to IV

* pokesearch config gen latest -> change default gen to the latest gen`,
	ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "latest"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gen := args[0]
		genNumeral := util.GenToNumeral(gen)
		viper.Set("gen", genNumeral)
		viper.WriteConfig()
		fmt.Printf(color.RedString("Gen set to %s (%s)!\n"), genNumeral, gen)
	},
}

func init() {
	configCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
