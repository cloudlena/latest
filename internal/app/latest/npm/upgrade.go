package npm

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/pkg/errors"
)

// Upgrade updates and upgrades all globally installed npm packages.
func (u Upgrader) Upgrade(upgrades chan<- latest.Upgrade) error {
	cmd := exec.Command("npm", "update", "-g")
	if u.verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	out, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "error running npm update -g")
	}

	npmUpgrades := upgradesFromOutput(string(out))
	for _, u := range npmUpgrades {
		upgrades <- u
	}

	return nil
}

func upgradesFromOutput(out string) []latest.Upgrade {
	re := regexp.MustCompile(`^\+ (.*)@(.*)$`)
	lines := strings.Split(out, "\n")
	upgrades := []latest.Upgrade{}
	for _, l := range lines {
		res := re.FindAllStringSubmatch(l, -1)
		if len(res) != 0 {
			u := latest.Upgrade{
				Upgrader:  name,
				Package:   res[0][1],
				VersionTo: res[0][2],
			}
			upgrades = append(upgrades, u)
		}
	}

	return upgrades
}
