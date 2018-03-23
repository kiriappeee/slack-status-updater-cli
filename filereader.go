package main

import (
	"os"
	"io/ioutil"

	"github.com/kiriappeee/ssucore"
)

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