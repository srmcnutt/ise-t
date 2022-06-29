package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getEnv() {
	var ise_pan string = os.Getenv("ISE_PAN")
	var ise_user string = os.Getenv("ISE_USER")
	var ise_password string = os.Getenv("ISE_PASSWORD")

	if ise_pan == "" {
		fmt.Println("ISE PAN Environment Variable missing!")
		os.Exit(1)
	}

	if ise_user == "" {
		fmt.Println("ISE USER Environment Variable missing!")
		os.Exit(1)
	}

	if ise_password == "" {
		fmt.Println("ISE Password Environment Variable missing!")
		os.Exit(1)
	}

}

func banner() {
	b, err := ioutil.ReadFile("art.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
	return
}

func main() {
	banner()
	getEnv()
}
