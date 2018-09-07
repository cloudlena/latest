package gem

// name is the name of this upgrader.
const name = "gem"

// Upgrader is the brew upgrader.
type Upgrader struct {
	verbose bool
}

// Make creates a new Upgrader.
func Make(verbose bool) Upgrader {
	return Upgrader{verbose: verbose}
}
