package ts_bbolt

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/containerd/containerd/v2/core/metadata/boltutil"
	bolt "go.etcd.io/bbolt"
)

func main1() {
	dbPath := flag.String("db", "", "path to metadata.db")
	root := flag.String("root", "", "snapshotter root dir")
	key := flag.String("key", "", "snapshot key/name")
	flag.Parse()

	if *dbPath == "" || *root == "" || *key == "" {
		fmt.Printf("usage: %s -db <metadata.db> -root <root> -key <key>\n",
			os.Args[0])
		os.Exit(2)
	}

	db, err := bolt.Open(*dbPath, 0444, &bolt.Options{ReadOnly: true})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id uint64
	err = db.View(func(tx *bolt.Tx) error {
		v1 := tx.Bucket([]byte("v1"))
		if v1 == nil {
			return fmt.Errorf("missing bucket v1")
		}
		ss := v1.Bucket([]byte("snapshots"))
		if ss == nil {
			return fmt.Errorf("missing bucket snapshots")
		}
		sbkt := ss.Bucket([]byte(*key))
		if sbkt == nil {
			return fmt.Errorf("snapshot key not found: %s", *key)
		}
		b := sbkt.Get([]byte("id"))
		if b == nil {
			return fmt.Errorf("id not found for key: %s", *key)
		}
		u, _ := binary.Uvarint(b)
		id = u
		return nil
	})
	if err != nil {
		panic(err)
	}

	idStr := strconv.FormatUint(id, 10)

	orig := filepath.Join(*root, "snapshots", idStr)
	// overlayfs 可能启用 short_base_paths，按 overlay 逻辑兼容：
	base := filepath.Dir(filepath.Dir(*root))
	short := filepath.Join(base, "l", idStr)

	path := orig
	if _, err := os.Stat(orig); err != nil {
		if _, err2 := os.Stat(short); err2 == nil {
			path = short
		}
	}

	fmt.Printf("id=%s\npath=%s\n", idStr, path)
}

func findByID(dbPath, idStr string) (*SnapInfo, error) {
	wantID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	db, err := bolt.Open(dbPath, 0444, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var out *SnapInfo
	err = db.View(func(tx *bolt.Tx) error {
		v1 := tx.Bucket([]byte("v1"))
		if v1 == nil {
			return fmt.Errorf("missing bucket v1")
		}
		ss := v1.Bucket([]byte("snapshots"))
		if ss == nil {
			return fmt.Errorf("missing bucket snapshots")
		}

		c := ss.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			sb := ss.Bucket(k)
			if sb == nil {
				continue
			}
			b := sb.Get([]byte("id"))
			if b == nil {
				continue
			}
			got, _ := binary.Uvarint(b)
			if got != wantID {
				continue
			}

			name := string(k)
			kind := readKind(sb)
			parent := string(sb.Get([]byte("parent")))
			var created, updated time.Time
			if err := boltutil.ReadTimestamps(sb, &created,
				&updated); err != nil {
				return err
			}
			labels, err := boltutil.ReadLabels(sb)
			if err != nil {
				return err
			}

			out = &SnapInfo{
				Name:    name,
				ID:      idStr,
				Kind:    byte(kind),
				Parent:  parent,
				Labels:  labels,
				Created: created,
				Updated: updated,
			}
			break
		}
		if out == nil {
			return fmt.Errorf("snapshot id not found: %s", idStr)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func main2() {
	dbPath := flag.String("db", "", "path to metadata.db")
	idStr := flag.String("id", "", "snapshot internal id")
	flag.Parse()

	if *dbPath == "" || *idStr == "" {
		fmt.Fprintf(os.Stderr,
			"usage: %s -db <metadata.db> -id <id>\n",
			os.Args[0],
		)
		os.Exit(2)
	}

	info, err := findByID(*dbPath, *idStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Printf("name=%s\nid=%s\nkind=%d\nparent=%s\n", info.Name,
		info.ID, info.Kind, info.Parent)
	fmt.Printf("created=%s\nupdated=%s\n", info.Created.Format(time.RFC3339),
		info.Updated.Format(time.RFC3339))
	fmt.Printf("labels=%v\n", info.Labels)
}
