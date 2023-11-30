package cache

import "sync"

type Cache struct {
	Mu       *sync.RWMutex
	UrlMap   map[string]string
	AliasMap map[string]string
}

func NewCache(mu *sync.RWMutex) Cache {
	return Cache{
		Mu:       mu,
		UrlMap:   make(map[string]string),
		AliasMap: make(map[string]string),
	}
}
