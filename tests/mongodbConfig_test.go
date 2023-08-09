package tests

import (
	"context"
	"project-p-back/pkg/config"
	"testing"
)

func TestInitMongoDb(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		expected bool
	}{
		{
			name:     "Valid MongoDB Connection",
			wantErr:  false,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.InitMongoDb()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitMongoDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			pingStatus := true // Assume the connection is successful
			if err := got.Ping(context.Background(), nil); err != nil {
				pingStatus = false
			}

			if pingStatus != tt.expected {
				t.Errorf("InitMongoDb() Ping status = %v, expected %v", pingStatus, tt.expected)
			}
		})
	}
}
