package mas

// name is the name of this upgrader.
const name = "mas"

// Upgrader is the mas upgrader.
type Upgrader struct {
	verbose bool
}

// New creates a new Upgrader.
func New(verbose bool) Upgrader {
	return Upgrader{verbose}
}
