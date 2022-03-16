package brew

// Upgrader is the brew upgrader.
type Upgrader struct {
	name    string
	verbose bool
}

// New creates a new Upgrader.
func New(verbose bool) *Upgrader {
	return &Upgrader{
		name:    "brew",
		verbose: verbose,
	}
}
