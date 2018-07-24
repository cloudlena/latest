package latest

// Upgrade is a successful upgrade.
type Upgrade struct {
	Upgrader    string
	Package     string
	VersionFrom string
	VersionTo   string
}

// Upgrader updates and upgrades stuff.
type Upgrader interface {
	Name() string
	Upgrade(upgradeChan chan<- Upgrade) error
}
