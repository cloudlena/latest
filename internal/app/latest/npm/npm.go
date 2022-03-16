package npm

// Upgrader is the brew upgrader.
type Upgrader struct {
	name    string
	verbose bool
}

// New creates a new upgrader.
func New(verbose bool) *Upgrader {
	return &Upgrader{
		name:    "npm",
		verbose: verbose,
	}
}
