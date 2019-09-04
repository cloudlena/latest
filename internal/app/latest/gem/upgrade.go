package gem

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
)

// upgradeRegex contains the name and version of upgrades.
var upgradeRegex = regexp.MustCompile(`^Successfully installed (.*)-(.*)$`)

// Upgrade updates and upgrades all globally installed gems.
func (u *upgrader) Upgrade(upgradesCh chan<- latest.Upgrade) error {
	cmd := exec.Command("gem", "update", "--force")
	if u.verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error running gem update: %w", err)
	}

	gemUpgrades := u.upgradesFromOutput(string(out))
	for i := range gemUpgrades {
		upgradesCh <- gemUpgrades[i]
	}

	cCmd := exec.Command("gem", "cleanup")
	if u.verbose {
		cCmd.Stdout = os.Stdout
		cCmd.Stderr = os.Stderr
	}

	err = cCmd.Run()
	if err != nil {
		return fmt.Errorf("error running gem cleanup: %w", err)
	}

	return nil
}

func (u *upgrader) upgradesFromOutput(out string) []latest.Upgrade {
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
