package gem

import (
	"testing"

	"github.com/mastertinner/latest/internal/app/latest"
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
			it: "parses gem output correctly",
			output: `
Updating installed gems
Updating bundler
Fetching: bundler-1.16.2.gem (100%)
Successfully installed bundler-1.16.2
Parsing documentation for bundler-1.16.2
Installing ri documentation for bundler-1.16.2
Installing darkfish documentation for bundler-1.16.2
Done installing documentation for bundler after 6 seconds
Parsing documentation for bundler-1.16.2
Done installing documentation for bundler after 2 seconds
Gems updated: bundler

			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "gem",
					Package:   "bundler",
					VersionTo: "1.16.2",
				},
			},
		},
		{
			it: "parses multiple packages correctly",
			output: `
			Updating installed gems
Updating bundler
Fetching: bundler-1.16.2.gem (100%)
Successfully installed bundler-1.16.2
Parsing documentation for bundler-1.16.2
Installing ri documentation for bundler-1.16.2
Installing darkfish documentation for bundler-1.16.2
Done installing documentation for bundler after 6 seconds
Parsing documentation for bundler-1.16.2
Done installing documentation for bundler after 2 seconds
Gems updated: bundler
Updating rack
Fetching: rack-2.0.5.gem (100%)
Successfully installed rack-2.0.5
Parsing documentation for rack-2.0.5
Installing ri documentation for rack-2.0.5
Installing darkfish documentation for rack-2.0.5
Done installing documentation for rack after 2 seconds
Parsing documentation for rack-2.0.5
Done installing documentation for rack after 0 seconds
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "gem",
					Package:   "bundler",
					VersionTo: "1.16.2",
				},
				{
					Upgrader:  "gem",
					Package:   "rack",
					VersionTo: "2.0.5",
				},
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			u := upgrader{name: "gem"}

			upgrades := u.upgradesFromOutput(tc.output)

			is.Equal(upgrades, tc.expectedUpgrades) // upgrades
		})
	}
}
