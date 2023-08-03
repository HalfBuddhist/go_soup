package ts_bbolt

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

type SnapInfo struct {
	Name    string
	ID      string
	Kind    byte
	Parent  string
	Labels  map[string]string
	Created time.Time
	Updated time.Time
	// Path    string
}

// kind 保持与 DB 中一致：单字节。
func readKind(bkt *bolt.Bucket) byte {
	k := bkt.Get([]byte("kind"))
	if len(k) == 1 {
		return k[0]
	}
	return 0
}
