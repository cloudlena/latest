package mas

import (
	"testing"

	"github.com/cloudlena/latest/internal/app/latest"
	"github.com/matryer/is"
)

func TestUpgradesFromOutput(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it               string
		output           string
		expectedUpgrades []latest.Upgrade
	}{
		{
			it: "parses mas output correctly",
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
			it: "parses mas output correctly for multiple packages",
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
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			u := Upgrader{name: "mas"}

			upgrades := u.upgradesFromOutput(tc.output)

			is.Equal(upgrades, tc.expectedUpgrades) // upgrades
		})
	}
}
