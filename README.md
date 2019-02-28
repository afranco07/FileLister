# Golang File Lister

A utility written in Go that lists all of the files within the directory you pass to it.

This utility starts a TCP server.

## Running the Application

1. Clone this repo
2. Run `go run main.go`
3. The application should have started and you shoud see `Starting server on port 8080 (localhost)...`
4. You can send strings using the `printf` command in the terminal i.e.:
```bash
printf "/home/alberto/Downloads" | nc localhost 8080
```
4. You should see a repsonse in the terminal