package mas

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/pkg/errors"
)

// upgradeRegex contains the name and version of upgrades.
var upgradeRegex = regexp.MustCompile(`^ ?(.*) \((.*)\)$`)

// Upgrade updates and upgrades brew.
func (u Upgrader) Upgrade(upgrades chan<- latest.Upgrade) error {
	cmd := exec.Command("mas", "upgrade")
	if u.verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	out, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "error running brew upgrade")
	}

	masUpgrades := upgradesFromOutput(string(out))
	for i := range masUpgrades {
		upgrades <- masUpgrades[i]
	}

	return nil
}

func upgradesFromOutput(out string) []latest.Upgrade {
	lines := strings.Split(out, "\n")
	upgrades := []latest.Upgrade{}
	for _, l := range lines {
		for _, p := range strings.Split(l, ",") {
			res := upgradeRegex.FindAllStringSubmatch(p, -1)
			if len(res) != 0 {
				u := latest.Upgrade{
					Upgrader:  name,
					Package:   res[0][1],
					VersionTo: res[0][2],
				}
				upgrades = append(upgrades, u)
			}
		}
	}

	return upgrades
}
