package url

import (
	"net/url"
	"reflect"

	qlang "qlang.io/spec"
)

func init() {
	Exports["Values"] = qlang.NewType(reflect.TypeOf((*Values)(nil)).Elem())
	Exports["values"] = newValues
	Exports["parseQuery"] = parseQuery
}

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values url.Values

func newValues() *Values {
	return &Values{}
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v *Values) Get(key string) string {
	if v == nil || *v == nil {
		return ""
	}
	vs, ok := (*v)[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (v *Values) Set(key, value string) {
	if v != nil {
		(*v)[key] = []string{value}
	}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v *Values) Add(key, value string) {
	if v != nil {
		(*v)[key] = append((*v)[key], value)
	}
}

// Del deletes the values associated with key.
func (v *Values) Del(key string) {
	if v != nil {
		delete(*v, key)
	}
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (v *Values) Encode() string {
	if v != nil {
		return (*url.Values)(v).Encode()
	}
	return ""
}

// ParseQuery parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
func parseQuery(query string) (Values, error) {
	m, err := url.ParseQuery(query)
	return Values(m), err
}
