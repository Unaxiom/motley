package copy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Dir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, but destination directory must *not* exist.
// Symlinks are ignored and skipped.
func Dir(srcDir string, dstDir string) (err error) {
	srcDir = filepath.Clean(srcDir)
	dstDir = filepath.Clean(dstDir)

	fileInfo, err := os.Stat(srcDir)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dstDir)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dstDir, fileInfo.Mode())
	if err != nil {
		return
	}

	dirList, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return
	}

	for _, entry := range dirList {
		srcDirPath := filepath.Join(srcDir, entry.Name())
		dstDirPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			err = Dir(srcDirPath, dstDirPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = File(srcDirPath, dstDirPath)
			if err != nil {
				return
			}
		}
	}

	return
}
