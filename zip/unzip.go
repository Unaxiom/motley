package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip accepts the tarball file name and untars the file to a target destination (which already exists)
func Unzip(zipArchive, target string) error {
	reader, err := zip.OpenReader(zipArchive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.Mode().IsDir() {
			err := os.MkdirAll(path, file.Mode())
			if err != nil {
				return err
			}
		} else if file.Mode().IsRegular() {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()

			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			targetFile.Chmod(file.Mode())
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		}

	}

	return nil
}
