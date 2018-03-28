package main

import (
	"os"
	"io/ioutil"

	"github.com/kiriappeee/slack-status-updater-core"
)
const defaultStatusList = `
- statusName: lunch
  emoji: bowl_with_spoon
  statusText: Having lunch

- statusName: resting
  emoji: bath
  statusText: Resting

- statusName: deepwork
  emoji: hammer_and_wrench
  statusText: In Deep work mode

- statusName: driving
  emoji: car
  statusText: On the road

- statusName: ping
  emoji: bell
  statusText: Working and available for pings :)

- statusName: gym
  emoji: weight_lifter
  statusText: Out getting exercise at the gym!

- statusName: meeting
  emoji: ''
  statusText: In a meeting

- statusName: happy
  emoji: grin
  statusText: ''
`
func getStatusesFromFile(pathToStatusFile string) ([] ssucore.Status, error) {
	_, err := os.Stat(pathToStatusFile)
	if os.IsNotExist(err) {
		return nil, err
	} else {
		data, err := ioutil.ReadFile(pathToStatusFile)
		if err == nil {
			return ssucore.ConvertTextToStructArray(string(data))
		} else {
			return nil, err
		}
	}
}

func checkAndSetupConfigDirectory(directoryPath string) error {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, 1)
		if err != nil {
			return err
		}
	}
	_, err = os.Stat(directoryPath + "/statuses.yaml")
	if os.IsNotExist(err) {
		f, err := os.OpenFile(directoryPath + "/statuses.yaml", os.O_WRONLY|os.O_CREATE, 0664)
		if _, err = f.Write([]byte(defaultStatusList)); err != nil{
			return err
		} else {
			if err = f.Close(); err != nil{
				return err
			}
		}
	}

	_, err = os.Stat(directoryPath + "/tokenconfig")
	if os.IsNotExist(err){
		f, err := os.OpenFile(directoryPath + "/tokenconfig", os.O_WRONLY|os.O_CREATE, 0664)
		if err = f.Close(); err != nil{
			return err
		}
	}
	return nil
}