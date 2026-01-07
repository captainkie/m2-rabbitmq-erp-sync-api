package model

import (
	"testing"
)

func TestConnectionLogs_Fields(t *testing.T) {
	l := ConnectionLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestAddLogs_Fields(t *testing.T) {
	l := AddLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestUpdateLogs_Fields(t *testing.T) {
	l := UpdateLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestStockLogs_Fields(t *testing.T) {
	l := StockLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestStoreLogs_Fields(t *testing.T) {
	l := StoreLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestPostflagLogs_Fields(t *testing.T) {
	l := PostflagLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestImageLogs_Fields(t *testing.T) {
	l := ImageLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}

func TestDailysaleLogs_Fields(t *testing.T) {
	l := DailysaleLogs{}
	if l.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if l.OrderID != "" {
		t.Error("OrderID should be empty by default")
	}
	if !l.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if l.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if l.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}
