package brew

// Name returns the name of the Homebrew executable.
func (u *upgrader) Name() string {
	return u.name
}
