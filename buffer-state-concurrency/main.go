package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var buffer bytes.Buffer
var byteCount = 0

var mutex sync.Mutex
var useMutex = true
var concurrentOperations = 1000

func main() {
	for index, arg := range os.Args {
		switch arg {
		case "--ignore-mutex":
			useMutex = false // Use "go run -race main.go" to see if any data races are encountered
		case "--operations":
			var argError error
			concurrentOperations, argError = strconv.Atoi(os.Args[index+1])
			if argError != nil {
				log.Fatal(argError)
			}
		}
	}

	wg.Add(2)
	go launchWrites()
	go launchReads()
	wg.Wait()

	var finalBytesRead = len(buffer.Next(buffer.Len()))
	log.Println("post concurrency bytes read: " + fmt.Sprint(finalBytesRead))
	byteCount += finalBytesRead
	log.Println("final byte count: " + fmt.Sprint(byteCount))
}

func launchReads() {
	defer wg.Done()
	for i := 0; i < concurrentOperations; i++ {
		wg.Add(1)
		go doARead()
	}
}

func doARead() {
	defer wg.Done()
	if useMutex {
		mutex.Lock()
		defer mutex.Unlock()
	}
	bytesRead := len(buffer.Next(buffer.Len()))
	log.Println("read " + fmt.Sprint(bytesRead))
	byteCount += bytesRead
}

func launchWrites() {
	defer wg.Done()
	for i := 0; i < concurrentOperations; i++ {
		wg.Add(1)
		go doAWrite()
	}
}

func doAWrite() {
	defer wg.Done()
	if useMutex {
		mutex.Lock()
		defer mutex.Unlock()
	}
	_, err := buffer.Write([]byte{42})
	if err != nil {
		log.Fatal(err)
	}
}
