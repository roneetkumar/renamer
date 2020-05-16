package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	path, name string
}

func main() {

	dir := "sample"

	var toRename []file

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})

	for _, f := range toRename {

		fmt.Printf("%q\n", f)
	}

	for _, orig := range toRename {

		var n file
		var err error

		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching: ", orig.path, err.Error())
		}

		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)

		err = os.Rename(orig.path, n.path)

		if err != nil {
			fmt.Println("Error renaming: ", orig.path, err.Error())
		}
	}
}

//  match return the new file name, or err
func match(filename string) (string, error) {

	parts := strings.Split(filename, ".")
	ext := parts[len(parts)-1]

	temp := strings.Join(parts[0:len(parts)-1], ".")

	parts = strings.Split(temp, "_")

	name := strings.Join(parts[0:len(parts)-1], "_")

	number, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't matched our pattern", filename)
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}
