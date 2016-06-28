package main

import ( "net/http"
	 "log"
	 "io"
	 "fmt"
	 "time"
	 "encoding/json"
)

type response_struct struct {
  Hostname string   `json:"hostname"`
  Date     time.Time     `json:"date"`
  Response string   `json:"response"`
}

func handleError(w http.ResponseWriter, err error) {
  log.Print(err)
  w.WriteHeader(http.StatusInternalServerError)
  io.WriteString(w, fmt.Sprintf("Server failed to process your request: %v", err))
}

func createJsonResponse(msg string) (resp string, err error) {
  response := response_struct {Hostname: hostname, Date: time.Now().UTC(), Response: msg}
  parsed, err := json.Marshal(response)
  resp = string(parsed)
  return
}
