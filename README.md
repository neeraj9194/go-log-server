# go-log-server
A log server to receive logs from different clients and store them in S3

## Install

To install you can use makefile or build using commands

```
make build
OR
go build main.go
```

To run, 

```
make run OR ./go-log-server
```

To run tests
```
make tests
OR
go test ./... -v
```


## API DOCS 

The server application collects logs on REST API.

### POST Logs
URL: `/` 

METHOD: `POST`

REQUEST BODY SCHEMA: `application/json`

BODY:

```
{
    "host":"ABC-123",
    "level":"INFO",
    "timestamp":"2021-01-01T00:00:00Z",
    "message":"message...",
    "service":"generic",
    "http":
        {"url":"","client_ip":"","version":""}
}
```



### List Logs
URL: `/` 

METHOD: `GET`

PARAMS: (NOT DONE)

&nbsp;&nbsp;&nbsp;&nbsp;`search`: To search logs

&nbsp;&nbsp;&nbsp;&nbsp;`start_ts`: Logs from timestamp 

&nbsp;&nbsp;&nbsp;&nbsp;`end_ts`: Logs until timestamp

 
REQUEST BODY SCHEMA: `application/json`

BODY:

```
[{
    "host":"ABC-123",
    "level":"INFO",
    "timestamp":"2021-01-01T00:00:00Z",
    "message":"message...",
    "service":"generic",
    "http":
        {"url":"","client_ip":"","version":""}
}
...
]
```

