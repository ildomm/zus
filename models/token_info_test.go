package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"testing"
	"time"
)

// Test TokenInfo Validate method
func TestTokenInfoValidate(t *testing.T) { // Test name: TestTokenInfoValidate
	var tests = []struct { // Test table
		name string    // Test name
		m    TokenInfo // Test model
		want error     // Expected result
	}{ // Test cases
		{ // Test case
			name: "valid", // Test case name
			m: TokenInfo{ // Test model
				CreatedAt: strfmt.DateTime(time.Now()),            // Test model value
				ID:        "123e4567-e89b-12d3-a456-426614174000", // Test model value
			},
			want: nil, // Expected result
		},
		{
			name: "invalid id",
			m: TokenInfo{
				ID: "invalid",
			},
			want: errors.New(422, "validation failure list:\nid in body must be of type uuid: \"invalid\""),
		},
	}

	for _, tt := range tests { // Test loop
		t.Run(tt.name, func(t *testing.T) { // Test name: TestTokenInfoValidate/<test name>
			got := tt.m.Validate(nil) // Call method
			if got != nil {           // Test result
				if got.Error() != tt.want.Error() { // Test result
					t.Errorf("TokenInfo.Validate() error = %v, wantErr %v", got, tt.want) // Test error
				}
			}
		})
	}
}

// Test TokenInfo generating hash	// Test name: TestTokenInfoGeneratingHash
func TestTokenInfoGeneratingHash(t *testing.T) { // Test name: TestTokenInfoGeneratingHash
	var tests = []struct { // Test table
		name string    // Test name
		m    TokenInfo // Test model
		want string    // Expected result
	}{ // Test cases
		{ // Test case
			name: "valid", // Test case name
			m: TokenInfo{ // Test model
				CreatedAt: strfmt.DateTime(time.Now()),            // Test model value
				ID:        "123e4567-e89b-12d3-a456-426614174000", // Test model value
			},
			want: "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a",
		},
	}

	for _, tt := range tests { // Test loop
		t.Run(tt.name, func(t *testing.T) { // Test name: TestTokenInfoGeneratingHash/<test name>
			got := tt.m.GenerateHash() // Call method
			if got != tt.want {        // Test result
				t.Errorf("TokenInfo.GenerateHash() = %v, want %v", got, tt.want) // Test error
			}
		})
	}
}
