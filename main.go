package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// telepresence container environment
	TELEPRESENCE_MOUNTS = "TELEPRESENCE_MOUNTS"
	TELEPRESENCE_ROOT   = "TELEPRESENCE_ROOT"
)

func main() {
	fmt.Println("...Running...")
	defer fmt.Println("...End...")

	telepresenceRoot := os.Getenv(TELEPRESENCE_ROOT)
	telepresenceMounts := os.Getenv(TELEPRESENCE_MOUNTS)
	if telepresenceMounts == "" {
		return
	}

	for _, mount := range strings.Split(telepresenceMounts, ":") {
		if IsExist(mount) {
			if err := os.RemoveAll(mount); err != nil {
				fmt.Println(err)
				continue
			}
		}

		dirName := filepath.Dir(mount)
		if err := os.MkdirAll(dirName, 0755); err != nil {
			fmt.Println(err)
		}

		linkSrc := filepath.Join(telepresenceRoot, mount[1:])

		fmt.Println("...Add Symlink...")
		fmt.Println("link: ", linkSrc)
		fmt.Println("mount: ", mount)
		if err := os.Symlink(linkSrc, mount); err != nil {
			fmt.Println(err)
		}
	}
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// exist
		return true
	}
	// not exist
	return false
}
