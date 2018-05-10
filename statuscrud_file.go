package main

import (
	"github.com/kiriappeee/slack-status-updater-core"
)

type statusCRUDFileImplementation string

func (s statusCRUDFileImplementation) GetStatuses() []ssucore.Status {
	return []ssucore.Status{
		ssucore.Status{"test", "emojiA", "Status One"},
		ssucore.Status{"test1", "emojiB", "Status Two"},
		ssucore.Status{"test2", "emojiC", "Status Three"},
		ssucore.Status{"test3", "emojiD", "Status Four"},
	}
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
