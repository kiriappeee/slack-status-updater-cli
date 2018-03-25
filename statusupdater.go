package main

import (
	"errors"

	"github.com/kiriappeee/slack-status-updater-core"
)

func pickStatusFromList( statusName string, s []ssucore.Status) (ssucore.Status, error) {
	for i := 0; i < len(s); i++ {
		if s[i].StatusName == statusName {
			return s[i], nil
		}
	}
	return ssucore.Status{}, errors.New("Could not find the requested status")
}