package ts_os

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

// moveFileOrDir moves a file or directory from src to dst.
// If src is a file, it will be renamed/moved to dst.
// If src is a directory, it will be moved to dst (dst should not exist).
func moveFileOrDir(src, dst string) error {
	// Get the file info for src
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// If it's a regular file, rename it
	if !fileInfo.IsDir() {
		return os.Rename(src, dst)
	}

	// If it's a directory, create the destination directory if it doesn't exist
	// and move all contents from src to dst
	if err := os.MkdirAll(dst, fileInfo.Mode()); err != nil {
		return err
	}

	// Walk the source directory and move each file/directory to the destination
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the relative path from src to the current file/directory
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// Construct the destination path
		newPath := filepath.Join(dst, relPath)

		// If it's a directory and we're not at the root, create it
		if info.IsDir() && relPath != "." {
			if err := os.MkdirAll(newPath, info.Mode()); err != nil {
				return err
			}
			// No need to copy directory contents here, filepath.Walk will handle it
			return nil
		}

		// Rename the file/directory to the new location
		return os.Rename(path, newPath)
	})
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mv <source> <destination>")
		os.Exit(1)
	}

	src := os.Args[1]
	dst := os.Args[2]

	err := moveFileOrDir(src, dst)
	if err != nil {
		log.Fatalf("Error moving %s to %s: %v(%v)", src, dst, err, reflect.TypeOf(err))
	}

	fmt.Printf("Successfully moved %s to %s\n", src, dst)
}
