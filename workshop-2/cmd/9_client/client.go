package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9000", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
