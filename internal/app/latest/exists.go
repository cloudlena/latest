package latest

import "os/exec"

// CmdExists checks if a command is available.
func CmdExists(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name) //nolint:gosec
	err := cmd.Run()

	return err == nil
}
