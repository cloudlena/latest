package mas

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
)

// upgradeRegex contains the name and version of upgrades.
var upgradeRegex = regexp.MustCompile(`^ ?(.*) \((.*)\)$`)

// Upgrade updates and upgrades brew.
func (u *upgrader) Upgrade(upgradesCh chan<- latest.Upgrade) error {
	cmd := exec.Command("mas", "upgrade")
	if u.verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error running brew upgrade: %w", err)
	}

	masUpgrades := u.upgradesFromOutput(string(out))
	for i := range masUpgrades {
		upgradesCh <- masUpgrades[i]
	}

	return nil
}

func (u *upgrader) upgradesFromOutput(out string) []latest.Upgrade {
	upgrades := []latest.Upgrade{}

	lines := strings.Split(out, "\n")
	for _, l := range lines {
		for _, p := range strings.Split(l, ",") {
			res := upgradeRegex.FindAllStringSubmatch(p, -1)
			if len(res) != 0 {
				u := latest.Upgrade{
					Upgrader:  u.name,
					Package:   res[0][1],
					VersionTo: res[0][2],
				}
				upgrades = append(upgrades, u)
			}
		}
	}

	return upgrades
}
