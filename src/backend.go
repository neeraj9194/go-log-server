package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
