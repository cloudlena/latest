package brew

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
			description: "parses brew output correctly",
			output: `
==> Downloading https://homebrew.bintray.com/bottles/packer-1.2.4.high_sierra.bottle.tar.gz
Already downloaded: /Users/Tobi/Library/Caches/Homebrew/packer-1.2.4.high_sierra.bottle.tar.gz
==> Pouring packer-1.2.4.high_sierra.bottle.tar.gz
==> Caveats
zsh completions have been installed to:
  /usr/local/share/zsh/site-functions
==> Summary
🍺  /usr/local/Cellar/packer/1.2.4: 7 files, 91.2MB
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "brew",
					Package:   "packer",
					VersionTo: "1.2.4",
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

func TestUpgradesFromCaskOutput(t *testing.T) {
	cases := []struct {
		description      string
		output           string
		expectedUpgrades []latest.Upgrade
	}{
		{
			description: "parses brew output correctly",
			output: `
==> Options
Include auto-update (-a): true
Include latest (-f): false
==> Updating Homebrew
Already up-to-date.
==> Finding outdated apps
       Cask                Current                                          Latest                                           A/U    Result
 1/18  alfred              3.6.1_910                                        3.6.1_910                                         Y   [   OK   ]
 2/18  appcleaner          3.4                                              3.4                                               Y   [   OK   ]
 3/18  docker              18.03.1-ce-mac65,24312                           18.03.1-ce-mac65,24312                            Y   [   OK   ]
 4/18  dropbox             latest                                           latest                                            Y   [   OK   ]
 5/18  google-chrome       67.0.3396.87                                     67.0.3396.87                                      Y   [   OK   ]
 6/18  graphiql            0.7.2                                            0.7.2                                                 [   OK   ]
 7/18  imageoptim          1.8.0                                            1.8.0                                             Y   [   OK   ]
 8/18  iterm2              3.1.6                                            3.1.6                                             Y   [   OK   ]
 9/18  macpass             0.7.3                                            0.7.3                                                 [   OK   ]
10/18  mongodb-compass     1.13.1                                           1.13.1                                                [   OK   ]
11/18  postman             6.1.3                                            6.1.3                                             Y   [   OK   ]
12/18  skype               8.23.0.10                                        8.23.0.10                                         Y   [   OK   ]
13/18  slack               3.2.0                                            3.2.0                                             Y   [   OK   ]
14/18  spotify             latest                                           latest                                            Y   [   OK   ]
15/18  sqlectron           1.29.0                                           1.29.0                                                [   OK   ]
16/18  virtualbox          5.2.12,122591                                    5.2.12,122591                                         [   OK   ]
17/18  visual-studio-code  1.24.0,6a6e02cef0f2122ee1469765b704faf5d0e0d859  1.24.1,24f62626b222e9a8313213fb64b10d741a326288   Y   [ FORCED ]
18/18  whatsapp            0.2.9737                                         0.2.9737                                          Y   [   OK   ]
==> Found outdated apps
     Cask                Current                                          Latest                                           A/U    Result
1/1  visual-studio-code  1.24.0,6a6e02cef0f2122ee1469765b704faf5d0e0d859  1.24.1,24f62626b222e9a8313213fb64b10d741a326288   Y   [ FORCED ]
==> Upgrading visual-studio-code to 1.24.1,24f62626b222e9a8313213fb64b10d741a326288
==> Satisfying dependencies
==> Downloading https://az764295.vo.msecnd.net/stable/24f62626b222e9a8313213fb64b10d741a326288/VSCode-darwin-stable.zip
######################################################################## 100.0%
==> Verifying checksum for Cask visual-studio-code
==> Installing Cask visual-studio-code
Warning: It seems there is already an App at '/Applications/Visual Studio Code.app'; overwriting.
==> Removing App '/Applications/Visual Studio Code.app'.
==> Moving App 'Visual Studio Code.app' to '/Applications/Visual Studio Code.app'.
==> Linking Binary 'code' to '/usr/local/bin/code'.
🍺  visual-studio-code was successfully installed!
			`,
			expectedUpgrades: []latest.Upgrade{
				{
					Upgrader:  "brew",
					Package:   "visual-studio-code",
					VersionTo: "1.24.1,24f62626b222e9a8313213fb64b10d741a326288",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			is := is.New(t)
			upgrades := upgradesFromCaskOutput(tc.output)
			is.Equal(upgrades, tc.expectedUpgrades) // upgrades
		})
	}
}
