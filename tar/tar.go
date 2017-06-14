package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhoonb/archivex"
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

// BigTar accepts a source directory and tars it to the target file. Target file should not exist. This method can handle big source directories as well.
func BigTar(sourceDir, target string) error {
	allContents, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	tar := new(archivex.TarFile)
	err = tar.Create(target)
	if err != nil {
		return err
	}
	for _, f := range allContents {
		if f.Mode().IsDir() {
			err = tar.AddAll(filepath.Join(sourceDir, f.Name()), true)
			if err != nil {
				return nil
			}
		} else if f.Mode().IsRegular() {
			filename := filepath.Join(sourceDir, f.Name())
			err = tar.AddFile(filename)
			if err != nil {
				return nil
			}
		}
	}
	err = tar.Close()
	if err != nil {
		return err
	}
	return nil
}
