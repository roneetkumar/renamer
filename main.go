package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")

var replaceStr = "$2 - $1 - $3 of $4.$5"

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "is it a dry run")
	flag.Parse()

	walkDir := "sample"

	var toRename []string

	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		return nil
	})

	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newfilename, _ := match(filename)
		newPath := filepath.Join(dir, newfilename)
		fmt.Printf("mv %s => %s\n", oldPath, newfilename)

		if !dry {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming: ", oldPath, newPath, err.Error())
			}
		}
	}
}

type matchResult struct {
	base string
	ext  string
}

func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didn't matched our pattern", filename)

	}

	tmp := re.ReplaceAllString(filename, replaceStr)
	return tmp, nil
}
