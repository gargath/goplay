package main

import ( "io"
 	 "net/http"
	 "log"
	 "os"
)

var hostname string

func serve(w http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }
  query, _ := parseQuery(req)
  log.Printf("Request: %v", query)
  result, _ := evaluateQuery(query)
  response, _ := createJsonResponse(query, result)
  io.WriteString(w, response)
}

func main() {
  var err error
  hostname, err = os.Hostname()
  if err != nil {
    log.Panicf("Unable to get hostname %v", err)
  }
  http.HandleFunc("/v1/", serve)
  log.Fatal(http.ListenAndServe(":8877", nil))
}
