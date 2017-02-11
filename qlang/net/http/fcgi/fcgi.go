package fcgi

import (
	"net/http/fcgi"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/fcgi",

	"serve": fcgi.Serve,
}
