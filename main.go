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
	ls bool
	//certs bool
)

type NodeList struct {
	SearchResult struct {
		Total     int
		Resources []node
	}
}

type node struct {
	Id   string
	Name string
	Link struct {
		Rel  string
		Href string
	}
	Certs []Cert
}

type CertList struct {
	Response []Cert
}

type Cert struct {
	ID                        string `json:"id"`
	FriendlyName              string `json:"friendlyName"`
	SerialNumberDecimalFormat string `json:"serialNumberDecimalFormat"`
	IssuedTo                  string `json:"issuedTo"`
	IssuedBy                  string `json:"issuedBy"`
	ValidFrom                 string `json:"validFrom"`
	ExpirationDate            string `json:"expirationDate"`
	UsedBy                    string `json:"usedBy"`
	KeySize                   int    `json:"keySize"`
	GroupTag                  string `json:"groupTag"`
	SelfSigned                bool   `json:"selfSigned"`
	SignatureAlgorithm        string `json:"signatureAlgorithm"`
	PortalsUsingTheTag        string `json:"portalsUsingTheTag"`
	Sha256Fingerprint         string `json:"sha256Fingerprint"`
	Link                      struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"link"`
}

// environment vars for connecting to ISE
var ise = make(map[string]string)

// store api endpoints in a map for easy retrieval
var endPoints = make(map[string]string)

func main() {
	//var iseNodes nodes
	ise = getEnv()
	flag.Parse()
	banner()

	// generate urls for api calls and store them
	// endPoint = API endpoint
	// the idea is to build a map of all the api endpoints we need to call
	initEndPoints()

	// get list of nodes in deployment
	iseNodes := getNodes()
	//fmt.Println(iseNodes)

	// print list of nodes in deployment
	//fmt.Println("Total Nodes: ", nodes.SearchResult.Total)

	// for i := 0; i < len(nodes.SearchResult.Resources); i++ {
	// 	fmt.Println(nodes.SearchResult.Resources[i].Name)
	// }
	for i := range iseNodes {
		fmt.Println("found Node: ", iseNodes[i].Name)
	}
	getCertificate(iseNodes[1].Name, iseNodes)
}

func init() {
	flag.BoolVar(&ls, "ls", false, "Lists nodes in deployment")
	//flag.BoolVar(&certs, "certs", false, "Lists certificates for nodes in deployment")
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

// initalize endpoint url map
func initEndPoints() {
	endPoints["pan"] = fmt.Sprintf("https://%s/ers/config/node", ise["pan"])
	endPoints["systemCerts"] = fmt.Sprintf("https://%s/api/v1/certs/system-certificate/", ise["pan"])
	return
}

// generic function to make rest api call to ISE and pass the body back
func iseCall(url string) []byte {

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
	//fmt.Println(res)

	//response body
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	return b
}

// return the index for a node using it's name
func getNodeByName(s string, n []node) int {
	for i := range n {
		if s == n[i].Name {
			fmt.Println("Node located!")
			fmt.Println(n[i])
			fmt.Println("Returning index ", i)
			return i

		}
	}
	fmt.Println("No match found")
	panic("node not found in database")
	return 999
}

// enumerate all nodes in the deployment and build inital data structure
func getNodes() []node {
	//use our nodelist struct to store the response
	var nodelist NodeList

	// make the api call
	res := iseCall(endPoints["pan"])

	error := json.Unmarshal(res, &nodelist)
	if error != nil {
		log.Println(error)
	}

	// build a slice of node items using the node struct
	var x []node
	for i := range nodelist.SearchResult.Resources {

		x = append(x, nodelist.SearchResult.Resources[i])

	}

	return x
}

// get the system certificate for a node, append to node data structure, return
func getCertificate(s string, n []node) {

	// build out our request string
	url := endPoints["systemCerts"] + s

	// make the api call
	res := iseCall(url)

	// make api call
	var certs CertList
	error := json.Unmarshal(res, &certs)
	if error != nil {
		log.Println(error)
	}

	// get index and and graft the cert list on to the node record
	index := getNodeByName(s, n)
	n[index].Certs = certs.Response

	for _, v := range n[index].Certs {
		fmt.Printf("\n\n\n")
		fmt.Println(v)
	}

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
