package main

import ( "net/http"
	 "net/url"
	 "log"
	 "io"
	 "fmt"
	 "time"
	 "encoding/json"
	 "strings"
	 "errors"
	 "strconv"
	 "github.com/crsmithdev/goexpr"
)

type response_struct struct {
  Hostname string   	 `json:"hostname"`
  Date     time.Time     `json:"date"`
  Query	   string        `json:"query"`
  Response string   	 `json:"response"`
}

func handleError(w http.ResponseWriter, err error) {
  log.Print(err)
  w.WriteHeader(http.StatusInternalServerError)
  io.WriteString(w, fmt.Sprintf("Server failed to process your request: %v", err))
}

func createJsonResponse(query string, result string) (resp string, err error) {
  response := response_struct {Hostname: hostname, Date: time.Now().UTC(), Query: query, Response: result}
  parsed, err := json.Marshal(response)
  resp = string(parsed)
  return
}

func evaluateQuery(query string) (result string, err error) {
  parsed, perr := goexpr.Parse(query)
  if perr != nil {
    log.Printf("Failed to parse expression: %v", perr)
    err = perr
    return
  }
  answer, cerr := goexpr.Evaluate(parsed, nil)
  if cerr != nil {
    log.Printf("Failed to evaluate expression: %v", cerr)
    err = cerr
    return
  }
  result = strconv.FormatFloat(answer, 'f', 2, 64)
  return
}

func parseQuery(req *http.Request) (result string, err error) {
  uri, _ := url.QueryUnescape(req.RequestURI)
  if !strings.HasPrefix(uri, "/v1/") {
    err = errors.New("URI not a /v1/ query")
    log.Printf("Wrong query")
    return
  }
  result = uri[4:len(uri)]
  return
}
