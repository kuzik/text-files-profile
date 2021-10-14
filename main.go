package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

type fileProfile []int
type rezProfile struct {
	sum   int
	count int
}

/*
* The project idea: write a CLI tool for creating files profile.
* Loop over all folders and files, collect lengths of each file lines
* write report with average lines
 */
func main() {
	dir := os.Args[1]
	if dir == "" {
		panic("Missed required parameter")
	}

	profile, err := CollectFolderProfile(dir)
	if err != nil {
		return
	}

	ProcessProfiles(profile)
}

func CollectFolderProfile(dir string) ([]fileProfile, error) {
	fileProfiles := make([]fileProfile, 1)
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		if item.IsDir() {
			subFolderProfile, err := CollectFolderProfile(dir + "/" + item.Name())
			if err != nil {
				return nil, err
			}
			fileProfiles = append(fileProfiles, subFolderProfile...)
		} else {
			// handle file there
			fileProfile, err := CollectFileProfile(dir + "/" + item.Name())
			if err != nil {
				return nil, err
			}
			fileProfiles = append(fileProfiles, fileProfile)
		}
	}

	return fileProfiles, nil
}

func CollectFileProfile(filePath string) (fileProfile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileProfile := make(fileProfile, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileProfile = append(fileProfile, utf8.RuneCountInString(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileProfile, nil
}

func ProcessProfiles(fileProfiles []fileProfile) fileProfile {

	sort.Slice(fileProfiles, func(i, j int) bool {
		return len(fileProfiles[i]) > len(fileProfiles[j])
	})

	res := make([]rezProfile, len(fileProfiles[0]))
	for _, singleFileProfile := range fileProfiles {
		for row, length := range singleFileProfile {
			res[row].sum += length
			res[row].count++
		}
	}

	for i, row := range res {
		fmt.Printf("%v: %s\n", i, strings.Repeat("*", row.sum/row.count))
	}
	return nil
}
