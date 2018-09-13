package brew

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/pkg/errors"
)

// These regexes contain the name and version of upgrades.
var (
	upgradeRegex     = regexp.MustCompile(`^==> Pouring (.*)-(.*)\.[a-zA-Z_]*\.bottle\.tar\.gz$`)
	caskUpgradeRegex = regexp.MustCompile(`^==> Upgrading (.*) to (.*)$`)
)

// Upgrade updates and upgrades brew.
func (u *upgrader) Upgrade(upgradesCh chan<- latest.Upgrade) error {
	udCmd := exec.Command("brew", "update")
	if u.verbose {
		udCmd.Stdout = os.Stdout
		udCmd.Stderr = os.Stderr
	}
	err := udCmd.Run()
	if err != nil {
		return errors.Wrap(err, "error running brew update")
	}

	ugCmd := exec.Command("brew", "upgrade", "--cleanup")
	if u.verbose {
		ugCmd.Stdout = os.Stdout
		ugCmd.Stderr = os.Stderr
	}
	ugOut, err := ugCmd.Output()
	if err != nil {
		return errors.Wrap(err, "error running brew upgrade")
	}
	ugUpgrades := u.upgradesFromOutput(string(ugOut))
	for i := range ugUpgrades {
		upgradesCh <- ugUpgrades[i]
	}

	cuCmd := exec.Command("brew", "cu", "--all", "--yes", "--cleanup")
	if u.verbose {
		cuCmd.Stdout = os.Stdout
		cuCmd.Stderr = os.Stderr
	}
	cuOut, err := cuCmd.Output()
	if err != nil {
		return errors.Wrap(err, "error running brew upgrade")
	}

	cuUpgrades := u.upgradesFromCaskOutput(string(cuOut))
	for i := range cuUpgrades {
		upgradesCh <- cuUpgrades[i]
	}

	cCmd := exec.Command("brew", "cleanup")
	if u.verbose {
		cCmd.Stdout = os.Stdout
		cCmd.Stderr = os.Stderr
	}
	err = cCmd.Run()
	if err != nil {
		return errors.Wrap(err, "error running brew cleanup")
	}

	return nil
}

func (u *upgrader) upgradesFromOutput(out string) []latest.Upgrade {
	lines := strings.Split(out, "\n")
	upgrades := []latest.Upgrade{}
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

func (u *upgrader) upgradesFromCaskOutput(out string) []latest.Upgrade {
	lines := strings.Split(out, "\n")
	upgrades := []latest.Upgrade{}
	for _, l := range lines {
		res := caskUpgradeRegex.FindAllStringSubmatch(l, -1)
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
