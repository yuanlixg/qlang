// +build go1.5

package reflect

import "reflect"

func init() {
	Exports["arrayOf"] = reflect.ArrayOf
	Exports["funcOf"] = reflect.FuncOf
}
