package tar

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Tar accepts a source directory and tars it to the target file. Target file should not exist.
func Tar(sourceDir, target string) error {
	// filename := filepath.Base(sourceDir)
	// target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	// fmt.Println("Filename is ", filename, " and target is ", target)
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(sourceDir)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(sourceDir)
	}

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, sourceDir))
		}

		if err := tarball.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarball, file)
		return err
	})
}
