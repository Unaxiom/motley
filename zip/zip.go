package zip

// import (
// 	// "io"
// 	"fmt"
// 	"io/ioutil"
// 	// "os"
// 	"path/filepath"
// 	// "strings"

// 	"github.com/jhoonb/archivex"
// )

// // Zip accepts a source directory and zips it to the target file. Target file should not exist.
// func Zip(sourceDir, target string) error {
// 	allContents, err := ioutil.ReadDir(sourceDir)
// 	if err != nil {
// 		return err
// 	}
// 	zip := new(archivex.ZipFile)
// 	err = zip.Create(target)
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Println("Source is ", sourceDir)
// 	// fmt.Println("Target is ", target)
// 	for _, f := range allContents {
// 		fmt.Println("File is ", f.Name(), " Dir: ", f.Mode().IsDir())
// 		if f.Mode().IsDir() {
// 			err = zip.AddAll(filepath.Join(sourceDir, f.Name()), true)
// 			if err != nil {
// 				return err
// 			}
// 		} else if f.Mode().IsRegular() {
// 			filename := filepath.Join(sourceDir, f.Name())
// 			// err = zip.AddFile(filename)
// 			err = zip.AddFileWithName(f.Name(), filename)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	err = zip.Close()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip accepts a source directory and zips it to the target file. Target file should not exist.
func Zip(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	var baseDir string
	info, err := os.Stat(source)
	if err != nil {
		return nil
	}
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	// fmt.Println("Base Dir is: ", baseDir)

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		// fmt.Println("Source is ", source)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			// header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			header.Name = strings.TrimPrefix(path, source)
		}
		// fmt.Println("Header is: ", header.Name)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
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
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
