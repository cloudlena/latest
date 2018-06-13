package brew

// name is the name of this upgrader.
const name = "brew"

// Upgrader is the brew upgrader.
type Upgrader struct {
	verbose bool
}

// New creates a new Upgrader.
func New(verbose bool) Upgrader {
	return Upgrader{verbose}
}