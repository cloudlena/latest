package mas

import (
	"testing"

	"github.com/mastertinner/latest/internal/app/latest"
	"github.com/matryer/is"
)

func TestUpgradesFromOutput(t *testing.T) {
	cases := []struct {
		description      string
		output           string
		expectedUpgrades []latest.Upgrade
	}{
		{
			description: "parses mas output correctly",
			output: `
Upgrading 1 outdated application:
Xcode (8.0)
==> Downloading Xcode
==> Installed Xcode
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "mas",
					Package:   "Xcode",
					VersionTo: "8.0",
				},
			},
		},
		{
			description: "parses mas output correctly for multiple packages",
			output: `
Upgrading 2 outdated applications:
Xcode (7.0), Screens VNC - Access Your Computer From Anywhere (3.6.7)
==> Downloading Xcode
==> Installed Xcode
==> Downloading iFlicks
==> Installed iFlicks
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "mas",
					Package:   "Xcode",
					VersionTo: "7.0",
				},
				{
					Upgrader:  "mas",
					Package:   "Screens VNC - Access Your Computer From Anywhere",
					VersionTo: "3.6.7",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			upgrades := upgradesFromOutput(tc.output)
			is.Equal(upgrades, tc.expectedUpgrades) // upgrades
		})
	}
}
