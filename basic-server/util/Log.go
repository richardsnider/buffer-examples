package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func Log(data interface{}) {
	logFilePath := os.Getenv("LOG_FILE_PATH")

	if logFilePath == "" {
		logFilePath = "/tmp/basic-server.json"
	}

	jsonBytes, marshalErr := json.Marshal(data)

	if marshalErr != nil {
		fmt.Println(marshalErr)
	}

	writeErr := os.WriteFile(logFilePath, jsonBytes, 0644)

	if writeErr != nil {
		fmt.Println(string(jsonBytes))
		fmt.Println(writeErr)
		return
	}

	logFileBytes, readErr := os.ReadFile(logFilePath)

	if readErr != nil {
		fmt.Println(readErr)
	}

	fmt.Println(logFileBytes)
}
