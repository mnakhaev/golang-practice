package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func stringMethods() {
	fmt.Println(
		strings.Contains("test_string", "test"), // true

		strings.Count("test_string", "t"), //3

		strings.HasPrefix("test_string", "test"), // true

		strings.HasSuffix("test_string", "string"), // true

		strings.Index("test_string", "t"), // 0

		strings.Join([]string{"a", "b", "c"}, "-"), // a-b-c

		strings.Split("a-b-c", "-"), //  [a b c]

		strings.Repeat("test", 3), // testtesttest

		strings.Replace("test_string", "test", "new", 1), // new_string

		strings.ToLower("TEST"), // test

		strings.ToUpper("test"), // TEST

	)
}

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
	//stringMethods()
	//readFile()
	//readFileShortly()
	createFile()
	readDirectory()
}
