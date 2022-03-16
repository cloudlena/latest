package mas

// Upgrader is the mas Upgrader.
type Upgrader struct {
	name    string
	verbose bool
}

// New creates a new upgrader.
func New(verbose bool) *Upgrader {
	return &Upgrader{
		name:    "mas",
		verbose: verbose,
	}
}
