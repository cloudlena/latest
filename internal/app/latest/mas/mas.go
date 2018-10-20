package mas

import "github.com/mastertinner/latest/internal/app/latest"

// upgrader is the mas upgrader.
type upgrader struct {
	name    string
	verbose bool
}

// New creates a new upgrader.
func New(verbose bool) latest.Upgrader {
	u := upgrader{
		name:    "mas",
		verbose: verbose,
	}
	return &u
}
