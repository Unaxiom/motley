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
		// fmt.Println("File is ", file.Name)
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(path, file.Mode())
			if err != nil {
				// fmt.Println("Couldn't create dir --> ", err.Error())
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		// targetFile, err := os.Create(path)
		if err != nil {
			return err
		}
		targetFile.Chmod(file.Mode())
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
