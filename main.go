package main

import (
	// "bytes"
	"fmt"
	// "os"
	"os/exec"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func subfolders( path string) []string {

	files, err := ioutil.ReadDir(path)
	pathM, _ := filepath.Abs(path)
    if err != nil {
        log.Fatal(err)
    }

	var subpath []string
    for _, file := range files {
        // fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			fileFullPath := filepath.Join(pathM, file.Name())
			subpath = append(subpath, fileFullPath)
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

func main () {
	
	rootPath := "/"
	rootFolders := subfolders(rootPath)
	
	remove(rootFolders, "/proc")
	remove(rootFolders, "/sys")
	
	var subpath2 []string
	for _, path := range rootFolders {
		for _, k := range subfolders(path) {
			subpath2 = append(subpath2, k)
		}
	}

	var subpath3 []string
	for _, path := range subpath2 {
		for _, k := range subfolders(path) {
			subpath3 = append(subpath3, k)
		}
	}
	
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

	// fmt.Println(m)
	return m

}

func diff(a , b map[string]int) []string {

	var s []string
	for k, v :=range a {
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