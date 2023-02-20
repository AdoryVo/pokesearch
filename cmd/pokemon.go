package cmd

import (
	"fmt"
	"net/http"
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

var pokemonCmd = &cobra.Command{
	Aliases: []string{"p"},
	Use:     "pokemon [pokémon]",
	Short:   "Search information for a specified Pokémon.",
	Long:    `Search information for a specified Pokémon. Information will be delivered in the form of links to Bulbapedia sections.`,
	Example: `
* pokesearch pokemon poliwag -l -> retrieve link for Poliwag's learnset from the current default config gen (IV if 'pokesearch config gen 4' was used)

* pokesearch pokemon kadabra -elst -g 5 -> retrieve links for Kadabra's evolution, Gen V learnset, stats, and type effectiveness`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		name = strings.ToLower(name)
		name = strings.Title(name)
		name = strings.ReplaceAll(name, " ", "_")

		// Check if the specified Pokemon exists
		resp, err := http.Head(fmt.Sprintf(strings.TrimSpace(BaseURI), name, ""))
		if err != nil {
			color.Red("Network connection too slow or unavailable - skipping validity check...")
		} else if resp.StatusCode == 404 {
			color.Red("The Pokémon you requested could not be found - you may have a typo!")
			return
		}
		defer resp.Body.Close()

		fmt.Printf(color.RedString("Pokémon: ")+BaseURI, name, "")

		if Evolution {
			suffix := "#Evolution"
			fmt.Printf(color.CyanString("Evolution: ")+BaseURI, name, suffix)
		}
		if Learnset {
			genString := viper.Get("gen")
			if genString == nil || Gen != -1 { // If gen was specified or no config gen exists
				genString = util.GenToNumeral(Gen)
			}

			suffix := ""
			switch genString {
			case "Latest":
				suffix = "#By_leveling_up"
			default:
				suffix = fmt.Sprintf("/Generation_%s_learnset", genString)
			}

			fmt.Printf(color.BlueString("Learnset: ")+BaseURI, name, suffix)
		}
		if Stats {
			suffix := "#Stats"
			fmt.Printf(color.GreenString("Stats: ")+BaseURI, name, suffix)
		}
		if TypeEffectiveness {
			suffix := "#Type_effectiveness"
			fmt.Printf(color.MagentaString("Type effectiveness: ")+BaseURI, name, suffix)
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
	pokemonCmd.Flags().IntVarP(&Gen, "gen", "g", -1, "specify gen for learnset (uses config gen by default)")
	pokemonCmd.Flags().BoolVarP(&Learnset, "learnset", "l", false, "search learnset")
	pokemonCmd.Flags().BoolVarP(&Stats, "stats", "s", false, "search stats")
	pokemonCmd.Flags().BoolVarP(&TypeEffectiveness, "type", "t", false, "search type effectiveness")
}
