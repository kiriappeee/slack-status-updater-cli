package main

import (
	"testing"

	"github.com/kiriappeee/slack-status-updater-core"
)

func TestStatusIsPickedFromListOfStatusesCorrectly(t *testing.T) {
	statuses, err := ssucore.ConvertTextToStructArray(dummyStatusList)
	if err != nil {
		t.Fatalf("Received unexpected error. Error message is %s", err.Error())
	}
	pickedStatus, err := pickStatusFromList("resting", statuses)
	if err != nil {
		t.Fatalf("Received unexpected error. Error message is %s", err.Error())
	}
	if pickedStatus.StatusName != "resting" && pickedStatus.Emoji != "bath" && pickedStatus.StatusText != "Resting" {
		t.Fatalf("Picked status attributes were not correct. Received: \n\n Name: %s\nEmoji: %s\nText: %s\n", pickedStatus.StatusName, pickedStatus.Emoji, pickedStatus.StatusText)
	}
	pickedStatus, err = pickStatusFromList("not-real", statuses)
	if err == nil {
		t.Fatalf("Did not receive expected error when picking non existent status")
	}
	if err.Error() != "Could not find the requested status" {
		t.Fatalf("Did not receive correct error message. Received: %s.\n\n Was expecting: 'Could not find the requested status'", err.Error())
	}
}