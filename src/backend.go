package src

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
)

type FS struct {
	root string
}

func (fs FS) write(logs []LogStruct) {
	// Create soem kind of file pools to write efficiently

	for _, val := range logs {
		folderPath := path.Join(fs.root, val.Service)
		os.MkdirAll(folderPath, os.ModePerm)
		writeLog(path.Join(folderPath, val.Host), val.Message)
	}
}

func writeLog(path string, log string) {
	logLine := []byte(log)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	_, er := f.Write(logLine)
	f.WriteString("\n")
	check(er)
}

func (fs FS) readLogs(logsChannel chan string, service string, hostname string) {
	files := readDirectory(fs.root, service, hostname)
	for _, file := range files {
		readFile(file, logsChannel)
	}
	close(logsChannel)
}

func readDirectory(root string, service string, hostname string) []string {
	var files []string
	if hostname == "" {
		err := filepath.Walk(path.Join(root, service),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					files = append(files, path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	} else {
		f := path.Join(root, service, hostname)
		files = append(files, f)
	}

	return files
}

func readFile(srcFilePath string, logsChannel chan string) {

	file, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logsChannel <- scanner.Text()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
