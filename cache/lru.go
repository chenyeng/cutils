package cache

import (
	"github.com/hashicorp/golang-lru/v2/expirable"
	"time"
)

type lru[K comparable, V any] struct {
	*expirable.LRU[K, V]
}

func (l *lru[K, V]) Get(k K) (V, bool) {
	return l.LRU.Get(k)
}

func NewLruCache[K comparable, V any](cap int, ttl time.Duration) Cache[K, V] {
	return &lru[K, V]{LRU: expirable.NewLRU[K, V](cap, nil, ttl)}
}

func (l *lru[K, V]) Set(k K, v V) {
	_ = l.LRU.Add(k, v)
}
