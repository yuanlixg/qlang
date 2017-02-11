// +build go1.7

package context

import "context"

func init() {
	Exports["Canceled"] = context.Canceled
	Exports["DeadlineExceeded"] = context.DeadlineExceeded
	Exports["background"] = context.Background
	Exports["TODO"] = context.TODO
	Exports["withCancel"] = context.WithCancel
	Exports["withDeadline"] = context.WithDeadline
	Exports["withTimeout"] = context.WithTimeout
	Exports["withValue"] = context.WithValue
}
