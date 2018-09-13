package mas

import "github.com/mastertinner/latest/internal/app/latest"

// upgrader is the mas upgrader.
type upgrader struct {
	name    string
	verbose bool
}

// Make creates a new upgrader.
func Make(verbose bool) latest.Upgrader {
	u := upgrader{
		name:    "mas",
		verbose: verbose,
	}
	return &u
}
