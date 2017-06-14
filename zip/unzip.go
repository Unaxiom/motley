package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip accepts the tarball file name and untars the file to a target destination (which already exists)
func Unzip(zipArchive, target string) error {
	rdr, err := zip.OpenReader(zipArchive)
	if err != nil {
		return err
	}

	// if err := os.MkdirAll(target, 0755); err != nil {
	// 	return err
	// }

	for _, file := range rdr.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {

			if fileReader != nil {
				fileReader.Close()
			}

			return err
		}

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fileReader.Close()

			if targetFile != nil {
				targetFile.Close()
			}

			return err
		}

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			fileReader.Close()
			targetFile.Close()

			return err
		}

		fileReader.Close()
		targetFile.Close()
	}

	return nil
}
