package main

import (
	"crypto/tls"
	"encoding/json"
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

type NodeList struct {
	SearchResult struct {
		Total     int
		Resources []struct {
			Id   string
			Name string
			Link struct {
				Rel  string
				Href string
				Type string
			}
		}
	}
}

// environment vars for connecting to ISE
var ise = make(map[string]string)

// store api endpoints in a map for easy retrieval
var endPoints = make(map[string]string)

func main() {
	var nodelist NodeList
	ise = getEnv()
	flag.Parse()
	banner()

	// generate urls for api calls and store them
	initEndPoints()

	// TODO generalize api call action
	r := iseCall(endPoints["nodes"], nodelist)

	// test output
	fmt.Printf("type is %T\n", r)
	fmt.Println(r.SearchResult.Resources[0].Name)
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

//initalize endpoint url map
func initEndPoints() {
	endPoints["nodes"] = fmt.Sprintf("https://%s/ers/config/node", ise["pan"])
	return
}

// generic function to make rest api call to ISE and pass the body back
func iseCall(url string, nodelist NodeList) NodeList {

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

	// build header
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	// add basic authentication to our header
	req.SetBasicAuth(ise["user"], ise["password"])

	// execute request & assign to res variable
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// dump the header
	fmt.Println(res)

	//mdump the response body
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	error := json.Unmarshal(b, &nodelist)
	if error != nil {
		log.Println(error)
	}
	//fmt.Println(j)
	return nodelist

	// fmt.Println(string(b))
	// err := json.Unmarshal(b, &m)

}

// enumerate all nodes in the deployment and build inital data structure
func getNodes() {
	// build url
	url := endPoints["nodes"]

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

	// build header
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	// add basic authentication to our header
	req.SetBasicAuth(ise["user"], ise["password"])

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
