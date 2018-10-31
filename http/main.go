package main

import (
	"io"
	"fmt"
	"net/http"
	"os(os.Stdout, resp.Body)(os.Stdout, resp.Body)"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)	
}
