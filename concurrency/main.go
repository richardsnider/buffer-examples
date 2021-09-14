package main

import (
	"log"
	"runtime"
	"strconv"
	"time"
)

func main() {
	defer log.Println(deferable())

	var channel1 = make(chan string)
	var channel2 = make(chan string)
	print("started")

	go subroutine(channel1)
	print("goroutine #1 initiated")

	go subroutine(channel2)
	print("goroutine #2 initiated")

	select {
	case channel1Result := <-channel1:
		print("channel #1 message received \"" + channel1Result + "\"")
	case channel2Result := <-channel2:
		print("channel #2 message received \"" + channel2Result + "\"")
		// default:
		// 	print("no activity . . .")
	}

	// <-channel1
	// <-channel2
	log.Println("Program completed.")
}

func deferable() int {
	log.Println("blaaaa")
	return 42
}

func subroutine(channel chan string) {
	print("started")
	time.Sleep(3 * time.Second)
	print("sending message")
	channel <- "hello"
	print("message sent")
	time.Sleep(3 * time.Second)
	print("finished") // This will not get the chance to execute
}

func print(message string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		var prefix = file + "(" + strconv.Itoa(line) + "): "
		log.Println(prefix + message)
	} else {
		log.Fatal("Logging error!")
	}
}
