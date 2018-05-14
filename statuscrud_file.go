package main

import (
	"errors"
	"github.com/kiriappeee/slack-status-updater-core"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type statusCRUDFileImplementation string

func (s statusCRUDFileImplementation) GetStatuses() []ssucore.Status {
	if s != "" {
		statuses, _ := getStatusesFromFile(string(s))
		return statuses
	}
	return []ssucore.Status{}
}

func (s statusCRUDFileImplementation) GetStatusByKey(name string) (ssucore.Status, error) {
	if name != "" {
		statuses := s.GetStatuses()
		for i := 0; i < len(statuses); i++ {
			if statuses[i].StatusName == name {
				return statuses[i], nil
			}
		}
		return ssucore.Status{}, errors.New("Could not find the requested status")
	}
	return ssucore.Status{}, errors.New("Name of status to search for cannot be empty")
}

func (s statusCRUDFileImplementation) AddNewStatus(status ssucore.Status) error {
	if s!= "" {
		statuses := s.GetStatuses()
		statuses = append(statuses, status)
		log.Println(statuses)
		data , err := yaml.Marshal(statuses)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(string(s), data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s statusCRUDFileImplementation) DeleteStatusByName(name string) error {
	if s!= "" {
		statuses := s.GetStatuses()
		var newStatusList []ssucore.Status
		for _, status := range(statuses) {
			if status.StatusName != name {
				newStatusList = append(newStatusList, status)
			}
		}
		data , err := yaml.Marshal(newStatusList)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(string(s), data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
