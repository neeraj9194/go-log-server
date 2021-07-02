package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"
	"bufio"
)

func TestReadFile(t *testing.T) {
	testString := `2009/01/23 01:23:23 Hello world this is a log message.`

	// Create a temp file
	file, err := ioutil.TempFile(".", "logfile")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	file.WriteString(fmt.Sprintf("%v\n", testString))

	testChannel := make(chan string)
	go readFile(file.Name(), testChannel)

	val := <-testChannel
	if val != testString {
		t.Fatal("Failed.")
	}
}

func TestReadDirectory(t *testing.T) {
	// Create a temp dir and files
	root_dir := "test_dir"
	os.MkdirAll(root_dir, os.ModePerm)
	file1, err := ioutil.TempFile(root_dir, "logfile1")
	if err != nil {
		log.Fatal(err)
	}
	file2, err2 := ioutil.TempFile(root_dir, "logfile2")
	if err2 != nil {
		log.Fatal(err2)
	}

	defer os.RemoveAll(root_dir)

	file_list := readDirectory(root_dir, "", "")
	if !reflect.DeepEqual(file_list, []string{file1.Name(), file2.Name()}) {
		t.Fatal("Failed.")
	}
}

func TestReadLogsFromDir(t *testing.T) {
	testString := `2009/01/23 01:23:23 Hello world this is a log message.`

	// Create a temp file
	root := "test_out"
	os.MkdirAll(root, os.ModePerm)
	file, err := ioutil.TempFile(root, "logfile")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(root)

	file.WriteString(fmt.Sprintf("%v\n", testString))

	testChannel := make(chan string)

	fs := FS{
		root,
	}
	go fs.readLogs(testChannel, "", "")

	val := <-testChannel
	if val != testString {
		t.Fatal("Failed.")
	}
}

func TestReadLogsFromFile(t *testing.T) {
	testString := `2009/01/23 01:23:23 Hello world this is a log message.`

	root := "test_out"
	service := "nginx"
	hostname := "neeraj"
	// Create a temp file
	os.MkdirAll(path.Join(root, service), os.ModePerm)

	file, err := ioutil.TempFile(path.Join(root, service), hostname)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(root)

	file.WriteString(fmt.Sprintf("%v\n", testString))

	testChannel := make(chan string)

	fs := FS{
		root,
	}
	go fs.readLogs(testChannel, service, filepath.Base(file.Name()))

	val := <-testChannel
	if val != testString {
		t.Fatal("Failed.")
	}
}


func TestWriteLog(t * testing.T) {
	root := "out_dir"
	logs := []LogStruct {
		{"neeraj", "Log line 1", "nginx"},
		{"neeraj", "Log line 2", "nginx"},
		{"neeraj", "Log line 1", "syslog"},
	}

	fs := FS{root}
	defer os.RemoveAll(root)
	fs.write(logs)

	// Read file1
	file, err := os.Open(path.Join(root,"nginx", "neeraj"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if(scanner.Text() != logs[i].Message) {
			log.Fatal("Failed")
		}
		i++
	}

	// Read file2
	file2, err := os.Open(path.Join(root, "syslog", "neeraj"))
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		if(scanner2.Text() != logs[i].Message) {
			log.Fatal("Failed")
		}
		i++
	}
} 
