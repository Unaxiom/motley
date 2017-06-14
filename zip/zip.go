package zip

import (
	// "io"
	"io/ioutil"
	// "os"
	"path/filepath"
	// "strings"

	"github.com/jhoonb/archivex"
)

// Zip accepts a source directory and zips it to the target file. Target file should not exist.
func Zip(sourceDir, target string) error {
	allContents, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	zip := new(archivex.ZipFile)
	err = zip.Create(target)
	if err != nil {
		return err
	}
	for _, f := range allContents {
		if f.Mode().IsDir() {
			err = zip.AddAll(filepath.Join(sourceDir, f.Name()), true)
			if err != nil {
				return nil
			}
		} else if f.Mode().IsRegular() {
			filename := filepath.Join(sourceDir, f.Name())
			err = zip.AddFile(filename)
			if err != nil {
				return nil
			}
		}
	}
	err = zip.Close()
	if err != nil {
		return err
	}
	return nil
}