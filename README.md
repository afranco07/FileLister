# Golang File Lister

![file-lister-demo](https://raw.githubusercontent.com/afranco07/gifImageStorage/master/golang_file_lister_demo.gif)

A utility written in Go that lists all of the files within the directory you pass to it.

This utility starts a TCP server.

## Running the Application

1. Download the binary under the `Releases` tab and execute it with `./golang_file_lister_1.0.0`
2. You can send strings using the `printf` command in the terminal i.e.:
```bash
printf "/usr/local/bin" | nc localhost 8080
```
3. You should see a response in the terminal

**OR**

1. Clone this repo
2. Run `go run main.go`
3. The application should have started and you shoud see `Starting server on port 8080 (localhost)...`
4. You can send strings using the `printf` command in the terminal i.e.:
```bash
printf "/usr/local/bin" | nc localhost 8080
```
4. You should see a response in the terminal

## Running Tests

To run the tests run:
```bash
go test
```

## Build Binary File

This file is already included under the `Releases` tab, but if you'd like to build it yourself you can do:
```bash
go build main.go
```