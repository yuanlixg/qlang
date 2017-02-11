package ioutil

import (
	"io/ioutil"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "io/ioutil",

	"Discard": ioutil.Discard,

	"nopCloser": ioutil.NopCloser,
	"readAll":   ioutil.ReadAll,
	"readDir":   ioutil.ReadDir,
	"readFile":  ioutil.ReadFile,
	"tempDir":   ioutil.TempDir,
	"tempFile":  ioutil.TempFile,
	"writeFile": ioutil.WriteFile,
}
