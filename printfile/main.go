package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(os.Args)
	myFile := os.Args[1]
	fmt.Println(myFile)

	b, err := ioutil.ReadFile(myFile)
	if err != nil {
		fmt.Print(err)
	}
	fs := string(b)

	fmt.Println(fs)
}
