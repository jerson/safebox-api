// Package cache ...
package cache

import (
	"github.com/allegro/bigcache"
	"time"
)

var client *bigcache.BigCache

// Connect ...
func Connect() (*bigcache.BigCache, error) {
	if client != nil {
		return client, nil
	}
	var err error
	client, err = bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		OnRemove:           nil,
	})

	return client, err
}
