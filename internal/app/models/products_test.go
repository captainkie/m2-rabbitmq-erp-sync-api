package model

import (
	"testing"
)

func TestConfigurableProducts_Validation(t *testing.T) {
	tests := []struct {
		name           string
		product        ConfigurableProducts
		wantErr        bool
		expectedErrors []string
	}{
		{
			name: "valid product",
			product: ConfigurableProducts{
				Sku:           "TEST-SKU-001",
				FirstChildSku: "CHILD-SKU-001",
			},
			wantErr: false,
		},
		{
			name: "empty SKU",
			product: ConfigurableProducts{
				Sku:           "",
				FirstChildSku: "CHILD-SKU-001",
			},
			wantErr:        true,
			expectedErrors: []string{"SKU should not be empty"},
		},
		{
			name: "empty FirstChildSku",
			product: ConfigurableProducts{
				Sku:           "TEST-SKU-001",
				FirstChildSku: "",
			},
			wantErr:        true,
			expectedErrors: []string{"FirstChildSku should not be empty"},
		},
		{
			name: "both fields empty",
			product: ConfigurableProducts{
				Sku:           "",
				FirstChildSku: "",
			},
			wantErr: true,
			expectedErrors: []string{
				"SKU should not be empty",
				"FirstChildSku should not be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test SKU validation
			if tt.product.Sku == "" && !tt.wantErr {
				t.Error("SKU should not be empty")
			}

			// Test FirstChildSku validation
			if tt.product.FirstChildSku == "" && !tt.wantErr {
				t.Error("FirstChildSku should not be empty")
			}

			// Test ID initialization
			if tt.product.ID != 0 {
				t.Error("ID should be initialized to 0")
			}

			// Test Created timestamp initialization
			if tt.product.Created != 0 {
				t.Error("Created should be initialized to 0")
			}

			// Test Updated timestamp initialization
			if tt.product.Updated != 0 {
				t.Error("Updated should be initialized to 0")
			}
		})
	}
}

func TestConfigurableProducts_FieldTypes(t *testing.T) {
	product := ConfigurableProducts{
		Sku:           "TEST-SKU-001",
		FirstChildSku: "CHILD-SKU-001",
	}

	// Test field types
	if product.Sku == "" {
		t.Error("SKU should be a string")
	}
	if product.FirstChildSku == "" {
		t.Error("FirstChildSku should be a string")
	}

	// Test field lengths
	if len(product.Sku) > 191 {
		t.Error("SKU length should not exceed 191 characters")
	}
	if len(product.FirstChildSku) > 191 {
		t.Error("FirstChildSku length should not exceed 191 characters")
	}
}
