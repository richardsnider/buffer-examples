package util

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func HttpHandler(responseWriter http.ResponseWriter, request *http.Request) {
	requestId := uuid.NewString()
	Log(map[string]interface{}{
		"requestTimestamp": time.Now().UTC().UnixNano(),
		"requestId ":       requestId,
		"method":           request.Method,
		"path":             request.URL.Path,
		"requestIpAddress": request.RemoteAddr,
	})

	responseWriter.Header().Add("Request-Id", requestId)
	randomString, err := RandString(10)

	if err != nil {
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.WriteHeader(200)
	responseWriter.Write([]byte(randomString))
}
