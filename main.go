package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := "birthday_001.txt"

	newName, err := match(filename, 4)

	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}

	fmt.Println(newName)

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
