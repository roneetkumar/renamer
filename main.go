package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	dir := "./sample"

	files, err := ioutil.ReadDir("./sample")

	if err != nil {
		panic(err)
	}

	count := 0

	var toRename []string

	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, origfilename := range toRename {

		orig := filepath.Join(dir, origfilename)
		newfilename, err := match(origfilename, count)
		if err != nil {
			panic(err)
		}

		newpath := filepath.Join(dir, newfilename)

		err = os.Rename(orig, newpath)
		fmt.Printf("mv %s => %s\n", orig, newpath)
		if err != nil {
			panic(err)
		}

	}
}

//  match return the new file name, or err
func match(filename string, total int) (string, error) {

	parts := strings.Split(filename, ".")
	ext := parts[len(parts)-1]

	temp := strings.Join(parts[0:len(parts)-1], ".")

	parts = strings.Split(temp, "_")

	name := strings.Join(parts[0:len(parts)-1], "_")

	number, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't matched our pattern", filename)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
