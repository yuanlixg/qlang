package signal

import (
	"os/signal"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/signal",

	"notify": signal.Notify,
	"stop":   signal.Stop,
}
