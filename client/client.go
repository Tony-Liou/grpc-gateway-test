package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://localhost:8081/v1/example/echo"

func main() {
	jsonStr := []byte(`{"value": "Say something"}`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response code: ", resp.StatusCode)
	fmt.Println("Response headers: ", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body: ", string(body))
}
