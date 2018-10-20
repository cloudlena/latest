package brew

import "github.com/mastertinner/latest/internal/app/latest"

// upgrader is the brew upgrader.
type upgrader struct {
	name    string
	verbose bool
}

// New creates a new Upgrader.
func New(verbose bool) latest.Upgrader {
	u := upgrader{
		name:    "brew",
		verbose: verbose,
	}
	return &u
}
