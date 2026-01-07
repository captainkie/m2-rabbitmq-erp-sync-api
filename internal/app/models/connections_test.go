package model

import (
	"testing"
)

func TestConnections_Fields(t *testing.T) {
	c := Connections{}
	if c.MessageCode != "" {
		t.Error("MessageCode should be empty by default")
	}
	if c.MessageDesc != "" {
		t.Error("MessageDesc should be empty by default")
	}
	if c.TotalRecordAdd != "0" && c.TotalRecordAdd != "" {
		t.Error("TotalRecordAdd should be '0' or empty by default")
	}
	if !c.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if c.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if c.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}
