package main

import (
	"github.com/kiriappeee/slack-status-updater-core"
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
	//stub
	return ssucore.Status{}, nil
}

func (s statusCRUDFileImplementation) AddNewStatus(status ssucore.Status) error {
	//stub
	return nil
}

func (s statusCRUDFileImplementation) DeleteStatusByName(name string) error {
	//stub
	return nil
}
