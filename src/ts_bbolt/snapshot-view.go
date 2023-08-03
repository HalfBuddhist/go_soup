package ts_bbolt

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/containerd/containerd/v2/core/metadata/boltutil"
	bolt "go.etcd.io/bbolt"
)

func main() {
	dbPath := flag.String("db", "/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/metadata.db", "path to metadata.db.")
	flag.Parse()
	if *dbPath == "" {
		fmt.Fprintf(os.Stderr,
			"usage: %s -db <metadata.db>\n",
			os.Args[0],
		)
		os.Exit(2)
	}

	db, err := bolt.Open(*dbPath, 0444, &bolt.Options{ReadOnly: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer db.Close()

	res := make([]*SnapInfo, 0)
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

			name := string(k)
			kind := readKind(sb)
			parent := string(sb.Get([]byte("parent")))
			var created, updated time.Time
			if err := boltutil.ReadTimestamps(sb, &created, &updated); err != nil {
				return err
			}
			labels, err := boltutil.ReadLabels(sb)
			if err != nil {
				return err
			}

			res = append(res, &SnapInfo{
				Name:    name,
				ID:      fmt.Sprintf("%d", got),
				Kind:    byte(kind),
				Parent:  parent,
				Labels:  labels,
				Created: created,
				Updated: updated,
			})
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("snapshot count: ", len(res))
	for _, v := range res {
		fmt.Printf("id: %s, name: %s, parent: %s, state: %v\n", v.ID, v.Name, v.Parent, v.Kind)
	}
}
