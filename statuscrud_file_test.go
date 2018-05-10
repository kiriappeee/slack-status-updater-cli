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
}
