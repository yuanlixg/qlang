// +build go1.5

package os

import "os"

func init() {
	Exports["lookupEnv"] = os.LookupEnv
}
