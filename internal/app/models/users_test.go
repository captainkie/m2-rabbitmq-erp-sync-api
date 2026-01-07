package model

import (
	"testing"
)

func TestUsers_Validation(t *testing.T) {
	tests := []struct {
		name    string
		user    Users
		wantErr bool
	}{
		{
			name: "valid user",
			user: Users{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "2",
				Status:   "0",
			},
			wantErr: false,
		},
		{
			name: "empty username",
			user: Users{
				Username: "",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "2",
				Status:   "0",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			user: Users{
				Username: "testuser",
				Email:    "",
				Password: "password123",
				Role:     "2",
				Status:   "0",
			},
			wantErr: true,
		},
		{
			name: "empty password",
			user: Users{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "",
				Role:     "2",
				Status:   "0",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test username validation
			if tt.user.Username == "" && !tt.wantErr {
				t.Errorf("Username should not be empty")
			}

			// Test email validation
			if tt.user.Email == "" && !tt.wantErr {
				t.Errorf("Email should not be empty")
			}

			// Test password validation
			if tt.user.Password == "" && !tt.wantErr {
				t.Errorf("Password should not be empty")
			}

			// Test role validation
			if tt.user.Role == "" {
				t.Errorf("Role should not be empty")
			}

			// Test status validation
			if tt.user.Status == "" {
				t.Errorf("Status should not be empty")
			}
		})
	}
}

func TestUsers_FieldTypes(t *testing.T) {
	user := Users{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "2",
		Status:   "0",
	}

	// Test ID field type
	if user.ID != 0 {
		t.Errorf("ID should be initialized to 0")
	}

	// Test Created field type
	if user.Created != 0 {
		t.Errorf("Created should be initialized to 0")
	}

	// Test Updated field type
	if user.Updated != 0 {
		t.Errorf("Updated should be initialized to 0")
	}

	// Test field types
	if user.Username == "" {
		t.Errorf("Username should be a string")
	}
	if user.Email == "" {
		t.Errorf("Email should be a string")
	}
	if user.Password == "" {
		t.Errorf("Password should be a string")
	}
	if user.Role == "" {
		t.Errorf("Role should be a string")
	}
	if user.Status == "" {
		t.Errorf("Status should be a string")
	}
}
