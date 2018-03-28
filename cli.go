package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"
	"io/ioutil"

	"github.com/urfave/cli"
	"github.com/kiriappeee/slack-status-updater-core"
)


func main(){
	if err := checkAndSetupConfigDirectory(os.Getenv("HOME") + "/.config/ssucli/"); err != nil {
		log.Fatal(err)
	}
	if err := checkAndGetToken(); err !=nil {
		log.Fatal(err)
	}
	app := cli.NewApp()
	app.Name = "Slack Status Updater"
	app.Usage = "Painless status updating via your Terminal"
	app.Version = "0.2.0"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "token",
			Usage: "Slack Legacy/API token to use with the CLI",
			FilePath: os.Getenv("HOME") + "/.config/ssucli/tokenconfig",
		},
	}
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
			Action: changeStatus,
		},
		{
			Name: "list",
			Usage: "Lists available statuses and their information",
			Action: listStatuses,
		},
	}
	app.Action = changeStatus
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func listStatuses(c *cli.Context) error {
	w := new(tabwriter.Writer)
	fileToReadStatusesFrom := os.Getenv("HOME") + "/.config/ssucli/statuses.yaml"
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	statusList, err := getStatusesFromFile(fileToReadStatusesFrom)
	if err != nil {
		fmt.Printf("There was an error while trying to read the status list: %s\n", err.Error())
	}
	fmt.Fprintln(w, "NAME", "\t", "EMOJI", "\t", "STATUS TEXT", "\t")
	
	for i := 0; i < len(statusList); i++ {
		fmt.Fprintln(w, statusList[i].StatusName, "\t", statusList[i].Emoji, "\t", statusList[i].StatusText, "\t")
	}
	w.Flush()
	return nil
}

func changeStatus(c *cli.Context) error{
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
}

func checkAndGetToken() error {
	_, err := os.Stat(os.Getenv("HOME") + "/.config/ssucli/tokenconfig")
	if os.IsNotExist(err) {
		return err
	} else {
		data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/ssucli/tokenconfig")
		if err != nil {
			return err
		} else {	
			if string(data) == "" {
				fmt.Print("Please enter the slack token value: ")
				fmt.Scanln(&data)
				f, err := os.OpenFile(os.Getenv("HOME") + "/.config/ssucli/tokenconfig", os.O_WRONLY, 0664); if err != nil {
					return err
				} else {
					if _, err = f.Write(data); err != nil{
						return err
					} else {
						if err = f.Close(); err != nil{
							return err
						}
					}
				}				
			}
		}
	}
	return nil
}