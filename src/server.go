package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LogStruct struct {
	// common attributes
	Host      string    `json:"host"`
	Level     string    `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Service   string    `json:"service"`

	// The structure can be extended to support diffrent types of services like HTTP, DB etc.
	Http HTTP `json:"http"`
}

type HTTP struct {
	URL      string `json:"url"`
	ClientIP string `json:"client_ip"`
	Version  string `json:"version"`
}

func logs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listLogs(w, r)
		return
	case "POST":
		storeLogs(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func storeLogs(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var logs []LogStruct
	err = json.Unmarshal(bodyBytes, &logs)

	// Write to file.
	fs := FS{
		"/home/neeraj/projects/go-log-server/test",
	}
	fs.write(logs)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func listLogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	logsChannel := make(chan LogStruct)
	fs := FS{
		"/home/neeraj/projects/go-log-server/test",
	}

	fs.readDir(logsChannel)

	var valList []LogStruct
	for {
		fmt.Println("hello2")
		select {
		case val := <-logsChannel:
			fmt.Println(val)
			valList = append(valList, val)
		default:
			json.NewEncoder(w).Encode(valList)
			return
		}
	}
}

func RunServer() {
	http.HandleFunc("/", logs)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
