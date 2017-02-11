// +build go1.5

package signal

import "os/signal"

func init() {
	Exports["ignore"] = signal.Ignore
	Exports["reset"] = signal.Reset
}
