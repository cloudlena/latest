package brew

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/pkg/errors"
)

// Upgrade updates and upgrades brew.
func (u Upgrader) Upgrade(upgrades chan<- latest.Upgrade) error {
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
	ugUpgrades := upgradesFromOutput(string(ugOut))
	for _, u := range ugUpgrades {
		upgrades <- u
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

	cuUpgrades := upgradesFromCaskOutput(string(cuOut))
	for _, u := range cuUpgrades {
		upgrades <- u
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

func upgradesFromOutput(out string) []latest.Upgrade {
	re := regexp.MustCompile(`^==> Pouring (.*)-(.*)\.[a-zA-Z_]*\.bottle\.tar\.gz$`)
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

func upgradesFromCaskOutput(out string) []latest.Upgrade {
	re := regexp.MustCompile(`^==> Upgrading (.*) to (.*)$`)
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
