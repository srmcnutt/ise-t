package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	ls    bool
	certs bool
)

func main() {
	flag.Parse()
	banner()
	getEnv()
}
func init() {
	flag.BoolVar(&ls, "ls", false, "Lists nodes in deployment")
	flag.BoolVar(&certs, "certs", false, "Lists certificates for nodes in deployment")

}

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
	fmt.Println(art)
	return
}

var art string = `

██╗███████╗███████╗      ████████╗
██║██╔════╝██╔════╝      ╚══██╔══╝
██║███████╗█████╗  █████╗   ██║   
██║╚════██║██╔══╝  ╚════╝   ██║   
██║███████║███████╗         ██║   
╚═╝╚══════╝╚══════╝         ╚═╝  v0.1

ISE Certificate toolbox 
 - by Steven McNutt, CCIE #6495. @densem0de on twitterz

`
