package main

import (
	"log"
	"os"
	"testing"
)
const dummyStatusList = `
- statusName: lunch
  emoji: chompy
  statusText: Having lunch

- statusName: resting
  emoji: bath
  statusText: Resting

- statusName: awesome
  emoji: awesome
  statusText: ''

- statusName: deep-work
  emoji: ''
  statusText: In Focus mode
`
func setUp(){
	homeDirectory := os.Getenv("HOME")
	_, err := os.Stat(homeDirectory+"/.config/ssuclitest")
	if os.IsNotExist(err){
		err := os.MkdirAll(homeDirectory+"/.config/ssuclitest", 1)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(homeDirectory+"/.config/ssuclitest/statuses.yaml", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		if _, err = f.Write([]byte(dummyStatusList)); err != nil {
			log.Fatal(err)
		}
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stat(homeDirectory+"/.config/ssuclitest/config")
	if os.IsNotExist(err){
		_, err := os.Create(homeDirectory+"/.config/ssuclitest/config")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func tearDown(){
	homeDirectory := os.Getenv("HOME")
	err := os.RemoveAll(homeDirectory+"/.config/ssuclitest")
	if err != nil {
		log.Fatal(err)
	}
}

func TestIfFileExists(t *testing.T){
	setUp()
	homeDirectory := os.Getenv("HOME")
	fileNameToSearchFor := homeDirectory + "/.config/ssuclitest/statuses.yaml"
	statusesToTest, err := getStatusesFromFile(fileNameToSearchFor)
	if err != nil {
		t.Fatalf("An error occurred. Expected no error. Error was %s", err.Error())
	}
	if len(statusesToTest) != 4 {
		t.Fatalf("Length of returned array was %d. Expected 4", len(statusesToTest))
	}
	tearDown()
}