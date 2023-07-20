package memory

import "sync"

type Registry struct {
	sync.RWMutex
}
