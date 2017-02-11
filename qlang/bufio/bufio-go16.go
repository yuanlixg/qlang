// +build go1.6

package bufio

import "bufio"

func init() {
	Exports["ErrFinalToken"] = bufio.ErrFinalToken
}
