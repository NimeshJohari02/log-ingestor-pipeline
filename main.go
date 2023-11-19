package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)
type LoggerRequestObject struct {
	Level       string    `json:"level"`
	Message     string    `json:"message"`
	ResourceID  string    `json:"resourceId"`
	Timestamp   time.Time `json:"timestamp"`
	TraceID     string    `json:"traceId"`
	SpanID      string    `json:"spanId"`
	Commit      string    `json:"commit"`
	Metadata    Metadata  `json:"metadata"`
}
type Metadata struct {
		ParentResourceID string `json:"parentResourceId"`
	}	
// The function `writeToFile` writes a LoggerRequestObject to a file in JSON format.
func writeToFile(req LoggerRequestObject) error {
	path := "./logs/" + "log_ingestor.log"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	// convert request to json
	content, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error while converting LoggerRequest to JSON ", err)
	}

	// convert json to string and write it to the file
	n, err := f.WriteString(string(content) + "\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %d bytes\n", n)
	f.Sync()

	return nil
}
// The handler function receives a POST request, decodes the request body into a LoggerRequestObject,
// and writes the object to a file.

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req LoggerRequestObject
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("Error while parsing request body ", err)
		}

		// write to file
		err = writeToFile(req)
		if err != nil {
			fmt.Println("Error while writing to file ", err)
		}
	}
}

// The main function sets up a server that listens on port 3000 and registers a handler function for
// the root route ("/").
func main() {
	http.HandleFunc("/", handler)            // Register the handler function for the root ("/") route
	err := http.ListenAndServe(":3000", nil) // Start the server on port 3000

	if err != nil {
		fmt.Println("Error:", err)
	}
}
