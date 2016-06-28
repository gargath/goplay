package main

import ( "io"
	 "io/ioutil"
 	 "net/http"
	 "log"
	 "encoding/json"
	 "os"
)

var hostname string

func serve(w http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    handleError(w, err)
    return
  }
  var jsonbody interface{}
  err = json.Unmarshal(body, &jsonbody)
  if err != nil {
    handleError(w, err)
    return
  }
  query, _ := parseQuery(req)
  log.Printf("Request: %v", query)
  response, _ := createJsonResponse(query)
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
