// +build go1.7

package exec

import "os/exec"

func init() {
	Exports["commandContext"] = exec.CommandContext
}
