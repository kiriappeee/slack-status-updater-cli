package main

import (
	"os"
	"testing"

	"github.com/kiriappeee/slack-status-updater-core"
)

func TestStatusesCanBeListed(t *testing.T) {
	fileSetUp()
	var sci ssucore.StatusCRUDInterface
	homeDirectory := os.Getenv("HOME")
	sci = statusCRUDFileImplementation(homeDirectory + "/.config/ssuclitest/statuses.yaml")
	statuses := ssucore.GetStatuses(sci)
	if len(statuses) != 4 {
		t.Fatalf("Did not receive the expected statuses")
	}
	fileTearDown()
}

func TestStatusCanBeRetrievedByKey(t *testing.T) {
	fileSetUp()
	var sci ssucore.StatusCRUDInterface
	homeDirectory := os.Getenv("HOME")
	sci = statusCRUDFileImplementation(homeDirectory + "/.config/ssuclitest/statuses.yaml")
	status, err := ssucore.GetStatusByKey("resting", sci)
	if err != nil {
		t.Fatalf("Received an error where non was expected. Received error was %s", err.Error())
	}
	if status.StatusName != "resting"{
		t.Fatalf("Did not received the expected status. Received %s", status)
	}

	status, err = ssucore.GetStatusByKey("notresting", sci)
	if err == nil {
		t.Fatalf("Did not receive an error when expecting one")
	}
	expectedError := "Could not find the requested status"
	receivedError := err.Error()

	if expectedError != receivedError {
		t.Fatalf("Did not receive the expected error. Received %s. Expected %s", receivedError, expectedError)
	}
	fileTearDown()
}

func TestStatusCanBeAdded(t *testing.T) {
	fileSetUp()
	var sci ssucore.StatusCRUDInterface
	homeDirectory := os.Getenv("HOME")
	sci = statusCRUDFileImplementation(homeDirectory + "/.config/ssuclitest/statuses.yaml")
	goodStatusToAdd := ssucore.Status{"mynewstatus", "emoji", "Status text"}
	err := ssucore.AddNewStatus(goodStatusToAdd, sci)
	if err != nil {
		t.Fatalf("Received an error where non was expected. Received error was: %s", err.Error())
	}
	status, _ := ssucore.GetStatusByKey("mynewstatus", sci)
	if status.StatusName != "mynewstatus" && status.Emoji != "emoji" && status.StatusText != "Status text" {
		t.Fatalf("Status retrieved after adding did not match expected values. Received %s", status)
	}
	fileTearDown()
}

func TestStatusCanBeDeleted(t *testing.T) {
	fileSetUp()
	var sci ssucore.StatusCRUDInterface
	homeDirectory := os.Getenv("HOME")
	sci = statusCRUDFileImplementation(homeDirectory + "/.config/ssuclitest/statuses.yaml")
	err := ssucore.DeleteStatusByName("resting", sci)
	if err != nil {
		t.Fatalf("Received an error where non was expected. Received error was: %s", err.Error())
	}

	_, err = ssucore.GetStatusByKey("resting", sci)
	if err == nil {
		t.Fatalf("Received no error. Was expecting: Could not find the requested status")
	}
	fileTearDown()
}
