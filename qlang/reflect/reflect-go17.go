// +build go1.7

package reflect

import "reflect"

func init() {
	Exports["structOf"] = reflect.StructOf
}
