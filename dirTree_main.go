package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func find_files(path string, pref string, output io.Writer, prOrNot bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var ref []os.FileInfo
	if !prOrNot {
		for i := 0; i < len(files); i++ {
			if files[i].IsDir() {
				ref = append(ref, files[i])
			}
		}
	} else {
		ref = files
	}
	for ind, file := range ref {
		printInner(file, pref, path, ind, len(ref), output, prOrNot)
	}
}

func printInner(file os.FileInfo, pref string, path string, ind int, leng int, output io.Writer, prOrNot bool) {
	var addit1 string
	if ind == leng-1 {
		addit1 = "└───"
	} else {
		addit1 = "├───"
	}
	if file.IsDir() {
		fmt.Fprintln(output, pref + addit1 + file.Name())
		pth := path + string(os.PathSeparator) + file.Name()
		if ind == leng-1 {
			find_files(pth, pref+"\t", output, prOrNot)
		} else {
			find_files(pth, pref+"│\t", output, prOrNot)
		}
	} else {
		var inf int64 = file.Size()
		var resstr string
		if inf > 0 {
			resstr = "(" + strconv.Itoa(int(inf)) + "b)"
		} else {
			resstr = "(empty)"
		}
		fmt.Fprintln(output, pref + addit1 + file.Name(), resstr)
	}
}

func dirTree(output io.Writer, path string, prOrNot bool) error {
	find_files(path, "", output, prOrNot)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
