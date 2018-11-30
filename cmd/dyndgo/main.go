package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ulm0/dyndgo/data"
	"github.com/ulm0/dyndgo/gmi"
)

const (
	dataFilename = "./data.yml"
)

func main() {
	var d data.Data
	_, err := d.ReadData(dataFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ip, err := gmi.GetIP()
	if err != nil {
		log.Fatalln("Error getting IP")
	}
	fmt.Printf("IP detected: %s\n", ip)

	err = d.UpdateDomains(ip)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
