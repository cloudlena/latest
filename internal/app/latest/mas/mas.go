package mas

// name is the name of this upgrader.
const name = "mas"

// Upgrader is the mas upgrader.
type Upgrader struct {
	verbose bool
}

// Make creates a new Upgrader.
func Make(verbose bool) Upgrader {
	return Upgrader{verbose: verbose}
}
