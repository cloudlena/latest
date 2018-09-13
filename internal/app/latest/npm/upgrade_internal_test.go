package npm

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
			description: "parses npm output correctly",
			output: `
/usr/local/bin/create-react-app -> /usr/local/lib/node_modules/create-react-app/index.js
+ create-react-app@1.5.2
added 5 packages from 3 contributors and updated 2 packages in 2.091s
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "npm",
					Package:   "create-react-app",
					VersionTo: "1.5.2",
				},
			},
		},
		{
			description: "parses multiple packages correctly",
			output: `
/usr/local/bin/create-react-app -> /usr/local/lib/node_modules/create-react-app/index.js
/usr/local/bin/tslint -> /usr/local/lib/node_modules/tslint/bin/tslint
npm WARN tslint@5.10.0 requires a peer of typescript@>=2.1.0 || >=2.1.0-dev || >=2.2.0-dev || >=2.3.0-dev || >=2.4.0-dev || >=2.5.0-dev || >=2.6.0-dev || >=2.7.0-dev || >=2.8.0-dev || >=2.9.0-dev but none is installed. You must install peer dependencies yourself.
npm WARN tsutils@2.27.1 requires a peer of typescript@>=2.1.0 || >=2.1.0-dev || >=2.2.0-dev || >=2.3.0-dev || >=2.4.0-dev || >=2.5.0-dev || >=2.6.0-dev || >=2.7.0-dev || >=2.8.0-dev || >=2.9.0-dev || >= 3.0.0-dev but none is installed. You must install peer dependencies yourself.

+ create-react-app@1.5.2
+ tslint@5.10.0
added 5 packages from 3 contributors and updated 3 packages in 2.446s
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "npm",
					Package:   "create-react-app",
					VersionTo: "1.5.2",
				},
				{
					Upgrader:  "npm",
					Package:   "tslint",
					VersionTo: "5.10.0",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			u := upgrader{name: "npm"}
			upgrades := u.upgradesFromOutput(tc.output)
			is.Equal(upgrades, tc.expectedUpgrades) // upgrades
		})
	}
}
