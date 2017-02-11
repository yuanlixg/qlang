package sync

import (
	"sync"
)

func init() {
	Exports["mutex"] = newMutex
	Exports["waitGroup"] = newWaitGroup
}

// -----------------------------------------------------------------------------

func newMutex() *sync.Mutex {
	return new(sync.Mutex)
}

func newWaitGroup() *sync.WaitGroup {
	return new(sync.WaitGroup)
}

// -----------------------------------------------------------------------------
