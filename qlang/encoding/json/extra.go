package json

import (
	"encoding/json"
	"strings"
	"syscall"
)

func init() {
	Exports["pretty"] = Pretty
	Exports["unMarshal"] = Unmarshal
}

// -----------------------------------------------------------------------------

// Pretty prints a value in pretty mode.
//
func Pretty(v interface{}) string {

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Unmarshal unmarshals a []byte or string
//
func Unmarshal(b interface{}) (v interface{}, err error) {

	switch in := b.(type) {
	case []byte:
		err = json.Unmarshal(in, &v)
	case string:
		err = json.NewDecoder(strings.NewReader(in)).Decode(&v)
	default:
		err = syscall.EINVAL
	}
	return
}

// -----------------------------------------------------------------------------
