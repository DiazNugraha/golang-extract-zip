package main

import (
	"archive/zip"
	"flag"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// parse args
	flag.Parse()
	zipFile := flag.Arg(0)

	// read zip file
	reader, _ := zip.OpenReader(zipFile)

	// extraction result directory
	destination := "my_example_unpacked"
	var filenames []string

	defer reader.Close()

	for _, file := range reader.File {
		fileExtraction(file, destination, &filenames)
	}
}

func fileExtraction(file *zip.File, destination string, filenames *[]string) bool {
	fp := filepath.Join(destination, file.Name)
	*filenames = append(*filenames, fp)
	if file.FileInfo().IsDir() {
		os.MkdirAll(fp, os.ModePerm)
		return true
	}
	os.MkdirAll(filepath.Dir(fp), os.ModePerm)
	outFile, _ := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	rc, _ := file.Open()
	io.Copy(outFile, rc)

	outFile.Close()
	rc.Close()
	return false
}
