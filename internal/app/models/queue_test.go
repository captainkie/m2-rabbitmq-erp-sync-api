package model

import "testing"

func TestConnectionQueues_RequiredFields(t *testing.T) {
	q := ConnectionQueues{}
	if q.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestAddQueues_RequiredFields(t *testing.T) {
	q := AddQueues{}
	if q.TransactionID != 0 {
		t.Error("TransactionID should be 0 by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestUpdateQueues_RequiredFields(t *testing.T) {
	q := UpdateQueues{}
	if q.TransactionID != 0 {
		t.Error("TransactionID should be 0 by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestStockQueues_RequiredFields(t *testing.T) {
	q := StockQueues{}
	if q.TransactionID != 0 {
		t.Error("TransactionID should be 0 by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestStoreQueues_RequiredFields(t *testing.T) {
	q := StoreQueues{}
	if q.TransactionID != 0 {
		t.Error("TransactionID should be 0 by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestDailysaleQueues_RequiredFields(t *testing.T) {
	q := DailysaleQueues{}
	if q.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if q.OrderID != "" {
		t.Error("OrderID should be empty by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestPostflagQueues_RequiredFields(t *testing.T) {
	q := PostflagQueues{}
	if q.TransactionID != 0 {
		t.Error("TransactionID should be 0 by default")
	}
	if q.JsonData != "" {
		t.Error("JsonData should be empty by default")
	}
}

func TestImageQueues_RequiredFields(t *testing.T) {
	q := ImageQueues{}
	if q.TransactionID != "" {
		t.Error("TransactionID should be empty by default")
	}
	if q.Image != "" {
		t.Error("Image should be empty by default")
	}
	if q.DirectoryPath != "" {
		t.Error("DirectoryPath should be empty by default")
	}
	if q.SyncDate != "" {
		t.Error("SyncDate should be empty by default")
	}
}
