package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ulm0/dyndgo/data"
	"github.com/ulm0/dyndgo/gmi"
	"gopkg.in/urfave/cli.v1"
)

const (
	dataFilename = "./data.yml"
)

func main() {
	app := cli.NewApp()
	app.Name = "dyndgo"
	app.Author = "ulm0"
	app.Email = "ulm0@innersea.xyz"
	app.Version = "0.1a"
	app.Usage = "Update yor DNS A records in DNSimple"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Value: "./data.yml",
			Usage: "YAML file with API token and A records",
		},
	}

	app.Action = func(c *cli.Context) error {
		var d data.Data
		_, err := d.ReadData(c.String("f"))
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
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
