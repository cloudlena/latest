package cmd

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/mastertinner/latest/internal/app/latest/brew"
	"github.com/mastertinner/latest/internal/app/latest/gem"
	"github.com/mastertinner/latest/internal/app/latest/mas"
	"github.com/mastertinner/latest/internal/app/latest/npm"
	"github.com/spf13/cobra"
)

const spinnerTimeout = 100 * time.Millisecond

// osCmd represents the os command.
var osCmd = &cobra.Command{
	Use:   "os",
	Short: "Update and upgrade your OS to the latest and greatest",
	Run: func(*cobra.Command, []string) {
		s := spinner.New(spinner.CharSets[14], spinnerTimeout)
		s.Start()
		upgraders := []latest.Upgrader{
			mas.New(verbose),
			brew.New(verbose),
			npm.New(verbose),
			gem.New(verbose),
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
			fmt.Printf(
				"%s: %s %s ==> %s\n",
				u.Upgrader,
				highlight.Sprint(u.Package),
				u.VersionFrom,
				highlight.Sprint(u.VersionTo),
			)
		}
		s.Stop()
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
			log.Fatal(fmt.Errorf("error upgrading %s packages: %w", name, err))
		}
	}
}
