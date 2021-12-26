package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/adoryvo/pokesearch/util"
)

var BaseURI string = "https://bulbapedia.bulbagarden.net/wiki/%s_(Pokémon)%s\n"
var Evolution bool
var Gen int
var Learnset bool
var Stats bool
var TypeEffectiveness bool

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		name = strings.ToLower(name)
		name = strings.Title(name)
		name = strings.ReplaceAll(name, " ", "_")

		if Evolution {
			suffix := "#Evolution"
			fmt.Printf(color.CyanString("Evolution: ") + BaseURI, name, suffix)
		}
		if Learnset {
			gen := viper.Get("gen")
			if Gen > 0{
				gen = util.GenToNumeral(strconv.Itoa(Gen))
			}
			suffix := ""

			switch gen {
			case "VIII":
				suffix = "#By_leveling_up"
			default:
				suffix = fmt.Sprintf("/Generation_%s_learnset", gen)
			}

			fmt.Printf(color.BlueString("Learnset: ") + BaseURI, name, suffix)
		}
		if Stats {
			suffix := "#Stats"
			fmt.Printf(color.GreenString("Stats: ") + BaseURI, name, suffix)
		}
		if TypeEffectiveness {
			suffix := "#Type_effectiveness"
			fmt.Printf(color.MagentaString("Type effectiveness: ") + BaseURI, name, suffix)
		}
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pokemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pokemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pokemonCmd.Flags().BoolVarP(&Evolution, "evolution", "e", false, "search evolution")
	pokemonCmd.Flags().IntVarP(&Gen, "gen", "g", -1, "specify gen")
	pokemonCmd.Flags().BoolVarP(&Learnset, "learnset", "l", false, "search learnset")
	pokemonCmd.Flags().BoolVarP(&Stats, "stats", "s", false, "search stats")
	pokemonCmd.Flags().BoolVarP(&TypeEffectiveness, "type", "t", false, "search type effectiveness")
}
