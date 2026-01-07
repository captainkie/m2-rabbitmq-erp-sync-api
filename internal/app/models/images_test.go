package model

import (
	"testing"
)

func TestImages_Fields(t *testing.T) {
	img := Images{}
	if img.Sku != "" {
		t.Error("Sku should be empty by default")
	}
	if img.ProductType != "" {
		t.Error("ProductType should be empty by default")
	}
	if !img.SyncDate.IsZero() {
		t.Error("SyncDate should be zero value by default")
	}
	if img.Image != "" {
		t.Error("Image should be empty by default")
	}
	if img.Created != 0 {
		t.Error("Created should be 0 by default")
	}
	if img.Updated != 0 {
		t.Error("Updated should be 0 by default")
	}
}
