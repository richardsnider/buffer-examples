package main

import (
	crand "crypto/rand"
	mrand "math/rand"
	"sort"
	"time"

	"encoding/binary"
	"fmt"
	"log"
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func main() {
	var src cryptoSource
	rand := mrand.New(src)

	var randomIntegers []int
	for i := 0; i < 100000; i++ {
		randomIntegers = append(randomIntegers, rand.Intn(100))
	}

	var startTime, endTime time.Time

	startTime = time.Now()
	sort.Ints(randomIntegers)
	endTime = time.Now()
	fmt.Println(endTime.Sub(startTime))

	startTime = time.Now()
	_ = countingSort(randomIntegers)
	endTime = time.Now()
	fmt.Println(endTime.Sub(startTime))

}

func countingSort(data []int) []int {
	var maxValue int
	for _, integer := range data {
		if integer > maxValue {
			maxValue = integer
		}
	}

	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(data)

	for i := 0; i < length; i++ {
		bucket[data[i]]++
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			data[sortedIndex] = j
			sortedIndex++
			bucket[j]--
		}
	}
	return data
}
