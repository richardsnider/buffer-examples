package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	var jsonBytes, marshalErr = json.Marshal(map[string]interface{}{
		"timestamp": time.Now().UTC().UnixMicro(),
		"user":      os.Getenv("USER"),
	})

	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	var writeErr = os.WriteFile("/tmp/data.json", jsonBytes, 0644)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	var fileBytes, readErr = os.ReadFile("/tmp/data.json")
	if readErr != nil {
		log.Fatal(readErr)
	}

	log.Println(string(fileBytes))
}
