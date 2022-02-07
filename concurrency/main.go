package main

import (
	"log"
	"runtime"
	"strconv"
	"time"
)

func print(message string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		var prefix = file + "(" + strconv.Itoa(line) + "): "
		log.Print(prefix + message)
	} else {
		log.Fatal("Logging error!")
	}
}

func subroutine(channel chan string, duration time.Duration, data string) {
	defer print("(subroutine) deferred end of function: \"" + data + "\"")
	print("(subroutine) start of function: \"" + data + "\"")
	time.Sleep(duration)
	channel <- data
	print("(subroutine) sent: \"" + data + "\"")
	time.Sleep(duration)
	// The end of this function will not be reached if the channel is not listened to
}

func main() {
	defer print("(main) deferred end of function")

	var channel1 = make(chan string)
	var channel2 = make(chan string)
	print("(main) start of function")

	go subroutine(channel1, 2*time.Second, "goroutine #1")
	print("(main) called for goroutine #1")

	go subroutine(channel2, 3*time.Second, "goroutine #2")
	print("(main) called for goroutine #2")

	channel1Result := <-channel1
	print("(main) channel #1 message received \"" + channel1Result + "\"")
	channel2Result := <-channel2
	print("(main) channel #2 message received \"" + channel2Result + "\"")

	// If we retrieve from either channel again we'll get an error
	// <-channel1
	// <-channel2

	time.Sleep(5 * time.Second)
	print("(main) end of function")
}
