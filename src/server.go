package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/neeraj9194/go-log-server/config"
)

type LogStruct struct {
	Host    string `json:"host"`
	Message string `json:"message"`
	Service string `json:"service"`
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
		config.RootDir,
	}
	fs.write(logs)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func listLogs(w http.ResponseWriter, r *http.Request) {
	logsChannel := make(chan string)
	
	// Prams
	query := r.URL.Query()
    service := query.Get("service")
    hostname := query.Get("hostname")
    
	fs := FS{
		config.RootDir,
	}

	go fs.readLogs(logsChannel, service, hostname)

	var valList []string
	for val := range logsChannel {
		valList = append(valList, val)
	}
	json.NewEncoder(w).Encode(valList)
}

func RunServer(serverPort string) {
	http.HandleFunc("/", logs)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	if err != nil {
		panic(err)
	}
}
