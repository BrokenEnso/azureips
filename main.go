package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ServiceTagSet struct {
	ChangeNumber int
	Cloud        string
	Values       []ServiceTag
}

type ServiceTag struct {
	Name       string
	Id         string
	Properties ServiceTagProperties
}
type ServiceTagProperties struct {
	ChangeNumber    int
	Region          string
	RegionId        int
	Platform        string
	SystemService   string
	AddressPrefixes []string
	NetworkFeatures []string
}

func (sts ServiceTagSet) FindTagByName(name string) (*ServiceTag, error) {
	for i := 0; i < len(sts.Values); i++ {
		if strings.Compare(sts.Values[i].Name, name) == 0 {
			return &sts.Values[i], nil
		}
	}
	return nil, errors.New("Now ServiceTag found by that name: " + name)
}

func main() {
	writeTo := flag.String("f", "stdout", "Path to file path to write Azure IPs to. Example -f /etc/host.allow.azure")
	serviceName := flag.String("s", "AzureCloud", "The name of the service tag to gt IPs for. Example -s AzureCloud.westus")
	flag.Parse()
	fmt.Println(*serviceName)

	var sendDataTo *os.File = os.Stdout
	var err error

	if strings.Compare(*writeTo, "stdout") != 0 {
		//if the file exists, it will be overwritten. use with caution kids
		sendDataTo, err = os.Create(*writeTo)
		if err != nil {
			log.Fatal(err)
		}
		defer sendDataTo.Close()
	}

	dlurl := getDlUrl()

	if len(dlurl) == 0 {
		log.Fatal("Couldn't find url for json file.")
	}

	resp, err := http.Get(dlurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cod ServiceTagSet

	err = json.Unmarshal(body, &cod)
	if err != nil {
		fmt.Println(err)
	}

	azureEast, err := cod.FindTagByName(*serviceName)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(azureEast.Properties.AddressPrefixes); i++ {
		sendDataTo.WriteString(azureEast.Properties.AddressPrefixes[i] + "\n")
	}
}

func getDlUrl() string {
	resp, err := http.Get("https://www.microsoft.com/en-us/download/confirmation.aspx?id=56519")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	finddlurl := regexp.MustCompile(`https:\/\/download\.microsoft\.com\/download\/[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*).json`)

	dlurl := finddlurl.Find(body)

	return string(dlurl)
}
