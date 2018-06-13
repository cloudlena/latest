package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/fatih/color"
	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/mastertinner/latest/internal/app/latest/brew"
	"github.com/mastertinner/latest/internal/app/latest/npm"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// osCmd represents the os command.
var osCmd = &cobra.Command{
	Use:   "os",
	Short: "Update and upgrade your OS to the latest and greatest",
	Run: func(cmd *cobra.Command, args []string) {
		upgraders := []latest.Upgrader{
			brew.New(verbose),
			npm.New(verbose),
		}

		upgrades := make(chan latest.Upgrade)
		var wg sync.WaitGroup
		wg.Add(len(upgraders))
		go func() {
			wg.Wait()
			close(upgrades)
		}()

		for _, u := range upgraders {
			go performUpgrades(u, &wg, upgrades)
		}

		highlight := color.New(color.FgGreen, color.Bold)
		for u := range upgrades {
			fmt.Printf("%s: %s %s ==> %s\n", u.Upgrader, highlight.Sprint(u.Package), u.VersionFrom, highlight.Sprint(u.VersionTo))
		}
	},
}

func init() {
	rootCmd.AddCommand(osCmd)
}

func performUpgrades(u latest.Upgrader, wg *sync.WaitGroup, upgrades chan<- latest.Upgrade) {
	defer wg.Done()
	name := u.Name()
	if latest.CmdExists(name) {
		err := u.Upgrade(upgrades)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error upgrading %s packages", name))
		}
	}
}
