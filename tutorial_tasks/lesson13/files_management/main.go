package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working dir:", dir)

	file, err := os.Open("tutorial_tasks/lesson13/requirements.txt")
	if err != nil {
		fmt.Println("Failed to open the file!")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func readFileShortly() {
	bs, err := ioutil.ReadFile("tutorial_tasks/lesson13/requirements.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func createFile() {
	file, err := os.Create("tutorial_tasks/lesson13/test_file.txt")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("This is test new file!")
}

func readDirectory() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func main() {
	readFile()
	readFileShortly()
	createFile()
	readDirectory()
}
