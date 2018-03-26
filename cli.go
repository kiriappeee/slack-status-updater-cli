package main

import (
	"os"
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"
	"github.com/kiriappeee/slack-status-updater-core"
)

func main(){
	app := cli.NewApp()
	app.Name = "Slack Status Updater"
	app.Usage = "Painless status updating via your Terminal"
	app.Version = "0.0.1 alpha super pre release"
	
	app.Commands = []cli.Command {
		{
			Name: "set",
			Usage: "Sets a status from a predefined list of statuses",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "token",
					Usage: "Slack Legacy/API token to use with the CLI",
					FilePath: os.Getenv("HOME") + "/.config/ssucli/tokenconfig",
				},
			},
			Action: func(c *cli.Context) error {
				fileToReadStatusesFrom := os.Getenv("HOME") + "/.config/ssucli/statuses.yaml"
				slackToken := strings.TrimSpace(c.String("token"))
				statusName := c.Args().Get(0)
				statusList, err := getStatusesFromFile(fileToReadStatusesFrom)
				if err != nil {
					fmt.Printf("There was an error while trying to read the status list: %s\n", err.Error())
					return nil
				}
				statusToSet, err := pickStatusFromList(statusName, statusList)
				if err != nil {
					fmt.Printf("There was an error while trying to find the status you named: %s\n", err.Error())
					return nil
				}
				res, err := statusToSet.SetMyStatus(ssucore.UpdateStatusViaSDK, slackToken)
				if res != "" {
					fmt.Printf("Status post result: %s\n", res)
				}
				if err != nil {
					fmt.Printf("There was an error while setting the status %s\n", err.Error())
					return nil
				}
				return nil
			},
		},
	}
	app.Action = func (c *cli.Context) error {
		fmt.Printf("Real token is %s", c.String("token"))
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}