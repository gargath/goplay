package main

import ( "io"
	 "io/ioutil"
 	 "net/http"
	 "log"
	 "encoding/json"
	 "os"
	 "fmt"
	 "time"
)

var hostname string

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
  response := response_struct {Hostname: hostname, Date: time.Now().UTC(), Response: fmt.Sprintf("Thank you from %v!\n", hostname)}
  response_json, _ := json.Marshal(response)
  io.WriteString(w, string(response_json))
}

func main() {
  var err error
  hostname, err = os.Hostname()
  if err != nil {
    log.Panicf("Unable to get hostname %v", err)
  }
  http.HandleFunc("/", serve)
  log.Fatal(http.ListenAndServe(":8877", nil))
}
