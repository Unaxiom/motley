package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Unaxiom/motley/zip"
)

func main() {
	localZip()
	// unzip()

}

func localZip() {
	cwd, _ := os.Getwd()
	fmt.Println("CWD is ", cwd)
	cwdList := strings.Split(cwd, "/")
	path := filepath.Join(strings.Join(cwdList[:len(cwdList)-1], "/"), "ProjectTemplate")
	fmt.Println(path)
	err := zip.Zip(path, filepath.Join(cwd, "ProjectTemplate.zip"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func unzip() {
	cwd, _ := os.Getwd()
	cwdList := strings.Split(cwd, "/")
	path := filepath.Join(strings.Join(cwdList[:len(cwdList)-1], "/"), "ProjectTemplate.zip")
	fmt.Println(path)
	// err := zip.Unzip(path, filepath.Join(cwd, "ProjectTemplate"))
	err := zip.Unzip(path, cwd)
	if err != nil {
		fmt.Println("Error is ", err.Error())
	}
}
