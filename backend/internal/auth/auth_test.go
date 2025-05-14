package auth

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		wantErr bool
	}{
		{
			name:    "valid token generation",
			userID:  1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateJWT(tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && token == "" {
				t.Error("GenerateJWT() returned empty token")
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		hashedPassword string
		want           bool
	}{
		{
			name:           "valid password",
			password:       "testPassword123",
			hashedPassword: "$2a$10$...", // Replace with actual hash
			want:           true,
		},
		{
			name:           "invalid password",
			password:       "wrongPassword",
			hashedPassword: "$2a$10$...", // Replace with actual hash
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.password, tt.hashedPassword)
			if got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
