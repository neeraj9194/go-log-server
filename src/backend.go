package src

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
)

type FS struct {
	root string
}

func (fs FS) write(logs []LogStruct) {
	// Create soem kind of file pools to write efficiently

	fmt.Println(logs)
	for _, val := range logs {
		writeLog(path.Join(fs.root, val.Service), val)
	}
}

func writeLog(path string, log LogStruct) {
	d, _ := json.Marshal(log)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	_, er := f.Write(d)
	f.WriteString("\n")
	check(er)
}

func (fs FS) readDir(logsChannel chan LogStruct) {
	var wg sync.WaitGroup
	x := []int{1}
	for range x {
		print("XXX")
		wg.Add(1)
		go readFile(&wg, "/home/neeraj/projects/go-log-server/test/generic", logsChannel)
	}
	wg.Wait()
}

func readFile(wg *sync.WaitGroup, srcFilePath string, logsChannel chan LogStruct) {
	defer wg.Done()

	file, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var log LogStruct
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)
		json.Unmarshal([]byte(x), &log)
		fmt.Println(log)
		logsChannel <- log
	}
	print("KJSADDSLKA\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
