package cache

import "sync"

type Cache struct {
	Mu       sync.RWMutex
	UrlMap   map[string]string
	AliasMap map[string]string
}
