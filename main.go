package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sync"
)

// WaitGroup for concurrency
var wg sync.WaitGroup

// ListFilesRecursively a function that takes a path/directory
// and prints the files in that path
// path: path that is being searched for
// relativePath: the current relative path respect to path
// connection: the current connected socket
func ListFilesRecursively(path string, relativePath string, connection io.Writer) {
	defer wg.Done()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(path, "does not exist")
		connection.Write([]byte(path + " does not exist\n"))
	}

	for _, file := range files {
		newPath := path + "/" + file.Name()
		relativePath := relativePath + "/" + file.Name()
		if file.IsDir() {
			wg.Add(1)
			go ListFilesRecursively(newPath, relativePath, connection)
		} else {
			fmt.Println(relativePath)
			connection.Write([]byte(relativePath + "\n"))
		}
	}
}

// ListFiles is a wrapper function for the recursive
// ListFilesRecursively function and adds concurency
func ListFiles(path string, connection io.Writer) {
	wg.Add(1)
	ListFilesRecursively(path, "", connection)
	wg.Wait()
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		requestLength, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error Reading:", err.Error())
			return
		}

		pathPassedIn := string(buffer[:requestLength])
		ListFiles(pathPassedIn, conn)
	}
}

func main() {
	fmt.Println("Starting server on port 8080 (localhost)...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("ERROR ON LISTNER", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error", err)
		}

		fmt.Println("connection", conn)
		go handleRequest(conn)
	}
}
