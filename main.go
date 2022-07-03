package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	ls    bool
	certs bool
)

var ise = make(map[string]string)

func main() {
	ise = getEnv()
	flag.Parse()
	banner()
	getNodes()
}

func init() {
	flag.BoolVar(&ls, "ls", false, "Lists nodes in deployment")
	flag.BoolVar(&certs, "certs", false, "Lists certificates for nodes in deployment")
}

// read in environment vars to connect to ISE
func getEnv() map[string]string {
	ise["pan"] = os.Getenv("ISE_PAN")
	ise["user"] = os.Getenv("ISE_USER")
	ise["password"] = os.Getenv("ISE_PASSWORD")

	if ise["pan"] == "" {
		fmt.Println("ISE PAN Environment Variable missing!")
		os.Exit(1)
	}

	if ise["user"] == "" {
		fmt.Println("ISE USER Environment Variable missing!")
		os.Exit(1)
	}

	if ise["password"] == "" {
		fmt.Println("ISE Password Environment Variable missing!")
		os.Exit(1)
	}

	return ise

}

// print banner describing application
func banner() {
	fmt.Println(art)
	return
}

// enumerate all nodes in the deployment and build iniital data structure
func getNodes() {
	// build url
	url := fmt.Sprintf("https://%s/ers/config/node", ise["pan"])

	// make a transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	// make a client
	client := &http.Client{Transport: tr}

	// set up request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(ise["user"], ise["password"])
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// execute request & assign to res variable
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// dump the header
	fmt.Println(res)

	//dump the response body
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}

// get the system certificates for each node and append to data structure
func getCertificates() {

}

// print a formatted list of system certificates
func printCertificateList() {

}

// expiration check

// export certificates

// generate ACDS certificates

// generate letsencrypt certificates

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
