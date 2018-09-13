package gem

import "github.com/mastertinner/latest/internal/app/latest"

// upgrader is the brew upgrader.
type upgrader struct {
	name    string
	verbose bool
}

// Make creates a new upgrader.
func Make(verbose bool) latest.Upgrader {
	u := upgrader{
		name:    "gem",
		verbose: verbose,
	}
	return &u
}
