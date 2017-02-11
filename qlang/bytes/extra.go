package bytes

import (
	"bytes"
	"fmt"
	"reflect"
)

func init() {
	Exports["buffer"] = newBuffer
	Exports["from"] = from
}

// -----------------------------------------------------------------------------

func newBuffer(a ...interface{}) *bytes.Buffer {

	if len(a) == 0 {
		return new(bytes.Buffer)
	}
	switch v := a[0].(type) {
	case []byte:
		return bytes.NewBuffer(v)
	case string:
		return bytes.NewBufferString(v)
	}
	panic("bytes.buffer() - unsupported argument type")
}

func from(v interface{}) []byte {

	switch args := v.(type) {
	case []int:
		b := make([]byte, len(args))
		for i, v := range args {
			b[i] = byte(v)
		}
		return b
	case string:
		return []byte(args)
	default:
		if v == nil {
			return nil
		}
		panic(fmt.Sprintf("can't convert from `%v` to []byte", reflect.TypeOf(v)))
	}
}

// -----------------------------------------------------------------------------
