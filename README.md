# go-log-server
A log server to receive logs from different clients and store them in S3

## Features

- [x] Server to accept Json logs
- [x] Store logs in file

## ToDo

- [x] Listing usign a API.
- [x] Accept unstructured logs as well.
- [ ] S3 integration using interfaces
- [ ] Tests
 

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

PARAMS: `service`, `hostname`

&nbsp;&nbsp;&nbsp;&nbsp;`search`: To search logs (Not done)

&nbsp;&nbsp;&nbsp;&nbsp;`service`: Filter logs based on service name 

&nbsp;&nbsp;&nbsp;&nbsp;`hostname`: Filter logs on hostname where it originated.

 
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


Server:
- It stores the logs, bucket them based on the identifier sent by the client.
- It should use the filesystem and S3 as backend storage.
- Support retrieval on exact text to begin with, later can support full text search.

