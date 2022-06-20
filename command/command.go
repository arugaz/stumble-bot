package command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ArugaZ/stumble-bot/app"
	"github.com/ArugaZ/stumble-bot/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCommand = &cobra.Command{
	Use:   "stumble",
	Short: "Stumble Bot for Push Rank",
	Long: `Stumble Bot for Push Rank:

	https://github.com/ArugaZ/stumble-bot`,
}

var runCommand = &cobra.Command{
	Use: "stumble",
	Run: func(cmd *cobra.Command, args []string) {
		if !strings.Contains(vars.Auth, "DeviceId") || !strings.Contains(vars.Auth, "Token") {
			fmt.Printf("%s[errors] %s%s\n%sgiven: %s", vars.ColorRed, vars.ColorReset, "invalid authorization", vars.ColorYellow, vars.Auth)
			return
		}
		if vars.Round < 0 || vars.Round > 3 {
			fmt.Printf("%s[errors] %s%s\n%sgiven: %d", vars.ColorRed, vars.ColorReset, "invalid round available 0-2", vars.ColorYellow, vars.Round)
			return
		}

		app.Run(&vars.Vars{
			Auth:  strings.Trim(vars.Auth, "\n"),
			Round: vars.Round,
			Url:   vars.Url,
		})
	},
}

func init() {
	auth, err := ioutil.ReadFile("auth.txt")
	if err != nil {
		fmt.Printf("%s[errors] %v", vars.ColorRed, err)
		return
	}
	cobra.OnInitialize()
	rootCommand.PersistentFlags().StringVarP(&vars.Auth, "auth", "a", string(auth), "Your Stumble Account Authorizaton")
	rootCommand.PersistentFlags().IntVarP(&vars.Round, "round", "r", vars.Round, "Round Type 0-2")
	viper.BindPFlag("Auth", rootCommand.Flags().Lookup("auth"))
	viper.BindPFlag("Round", rootCommand.Flags().Lookup("round"))
	rootCommand.AddCommand(runCommand)
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Printf("%s[errors] %v", vars.ColorRed, err)
		return
	}
}
