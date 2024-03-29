package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func subfolders(path string, depth int, subpath []string) []string {

	if depth > 0 {
		files, err := ioutil.ReadDir(path)
		pathM, _ := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		depth--
		for _, file := range files {
			if file.IsDir() {
				fileFullPath := filepath.Join(pathM, file.Name())
				subpath = append(subpath, fileFullPath)
				remove(subpath, "/proc")
				remove(subpath, "/mnt")
				if fileFullPath != "/proc" {
					subpath = subfolders(fileFullPath, depth, subpath)
				}
			}
		}
	}
	// fmt.Println(subpath)
	return subpath
}

func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func main() {

	//get parameters or set default
	var rootPath string
	if len(os.Args) > 1 {
		rootPath = os.Args[1]
	} else {
		rootPath = "/"
	}

	var array []string
	subpath3 := subfolders(rootPath, 3, array)
	fmt.Println(subpath3)

	//du -m -s /root
	state1 := tackPath(subpath3)
	time.Sleep(2 * time.Minute)
	state2 := tackPath(subpath3)
	diff(state1, state2)

}

func tackPath(path []string) map[string]int {
	m := make(map[string]int)
	for _, path := range path {
		output, err := exec.Command("du", "-ms", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		s := strings.Fields(string(output))[0]
		intVar, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		m[path] = intVar
	}
	return m
}

func diff(a, b map[string]int) []string {

	var s []string
	for k, v := range a {
		diff := v - b[k]
		if Abs(diff) > 5 {
			s = append(s, k)
		}
	}

	fmt.Println("++++++++++++++rusult:++++++++++++++++")
	fmt.Println(s)
	return s

}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
