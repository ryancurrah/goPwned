package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mavjs/goPwned"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "goPwned"
	app.Usage = "'Have I been pwned?' golang cli checker app"
	app.Version = gopwned.Version
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ye Myat Kaung (Maverick)",
			Email: "mavjs01@gmail.com",
		},
		cli.Author{
			Name:  "Chee Leong",
			Email: "klrkdekira@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "email",
			Aliases: []string{"m"},
			Usage:   "email address to look up all breaches associated with it",
			Action: func(c *cli.Context) {
				resp, err := gopwned.GetAllBreachesForAccount(c.Args().First(), "", "")
				if err != nil {
					log.Fatal(err)
				}

				b, err := json.MarshalIndent(resp, "", " ")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf(string(b))
			},
		},
	}

	app.Run(os.Args)
}
