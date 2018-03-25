package main

import (
	"os"
	"log"

	"github.com/urfave/cli"
	"github.com/kiriappeee/slack-status-updater-core"
)

func main(){
	app := cli.NewApp()
	slackToken := os.Getenv("SLACKPERSONALTOKEN")
	fileToReadStatusesFrom := os.Getenv("HOME") + "/.config/ssucli/statuses.yaml"
	app.Action = func(c *cli.Context) error {
		statusName := c.Args().Get(0)
		statusList, err := getStatusesFromFile(fileToReadStatusesFrom)
		if err != nil {
			log.Fatal(err)
		}
		statusToSet, err := pickStatusFromList(statusName, statusList)
		if err != nil {
			log.Fatal(err)
		}
		res, err := statusToSet.SetMyStatus(ssucore.UpdateStatusViaSDK, slackToken)
		if res != "" {
			log.Printf(res)
		}

		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}