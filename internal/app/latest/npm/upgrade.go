package npm

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/cloudlena/latest/internal/app/latest"
)

// upgradeRegex contains the name and version of upgrades.
var upgradeRegex = regexp.MustCompile(`^\+ (.*)@(.*)$`)

// Upgrade updates and upgrades all globally installed npm packages.
func (u *Upgrader) Upgrade(upgradesCh chan<- latest.Upgrade) error {
	cmd := exec.Command("npm", "update", "-g")
	if u.verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error running npm update -g: %w", err)
	}

	npmUpgrades := u.upgradesFromOutput(string(out))
	for i := range npmUpgrades {
		upgradesCh <- npmUpgrades[i]
	}

	return nil
}

func (u *Upgrader) upgradesFromOutput(out string) []latest.Upgrade {
	upgrades := []latest.Upgrade{}

	lines := strings.Split(out, "\n")
	for _, l := range lines {
		res := upgradeRegex.FindAllStringSubmatch(l, -1)
		if len(res) != 0 {
			u := latest.Upgrade{
				Upgrader:  u.name,
				Package:   res[0][1],
				VersionTo: res[0][2],
			}
			upgrades = append(upgrades, u)
		}
	}

	return upgrades
}
