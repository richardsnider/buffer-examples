package main

import (
	crand "crypto/rand"
	mrand "math/rand"
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

	var randomSortTargets []sortTarget
	for i := 0; i < 100; i++ {
		randomSortTargets = append(randomSortTargets, sortTarget{property2: rand.Float32()})
	}

	var startTime = time.Now()
	sortedTargets := getSortedVersion(randomSortTargets, func(target sortTarget) int { return int(target.property2 * float32(100)) })
	for _, target := range sortedTargets {
		fmt.Print(target.property2, ", ")
	}
	fmt.Println(time.Since(startTime))
}

type sortTarget struct {
	property1 int
	property2 float32
}

func getSortedVersion(data []sortTarget, evaluator func(target sortTarget) int) []sortTarget {
	var maxValue int
	for _, element := range data {
		var currentValue = evaluator(element)
		if currentValue > maxValue {
			maxValue = currentValue
		}
	}

	bucketLen := maxValue + 1
	bucket := make([][]sortTarget, bucketLen)

	sortedIndex := 0
	length := len(data)

	for i := 0; i < length; i++ {
		var bla = evaluator(data[i])
		bucket[bla] = append(bucket[bla], data[i])
	}

	for i := 0; i < bucketLen; i++ {
		for len(bucket[i]) > 0 {
			data[sortedIndex] = bucket[i][0]
			sortedIndex++
			bucket[i] = bucket[i][1:len(bucket[i])]
		}
	}
	return data
}
