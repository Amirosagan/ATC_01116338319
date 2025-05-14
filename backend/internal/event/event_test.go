package event

import (
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	tests := []struct {
		name        string
		event       Event
		wantErr     bool
		description string
	}{
		{
			name: "valid event",
			event: Event{
				Title:       "Test Event",
				Description: "Test Description",
				StartTime:   time.Now().Add(24 * time.Hour),
				EndTime:     time.Now().Add(48 * time.Hour),
				Location:    "Test Location",
				CreatorID:   1,
			},
			wantErr:     false,
			description: "Should create event successfully",
		},
		{
			name: "invalid time range",
			event: Event{
				Title:       "Test Event",
				Description: "Test Description",
				StartTime:   time.Now().Add(48 * time.Hour),
				EndTime:     time.Now().Add(24 * time.Hour),
				Location:    "Test Location",
				CreatorID:   1,
			},
			wantErr:     true,
			description: "Should fail when end time is before start time",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateEvent(&tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: CreateEvent() error = %v, wantErr %v", tt.description, err, tt.wantErr)
			}
		})
	}
}

func TestGetEvent(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *Event
		wantErr bool
	}{
		{
			name:    "existing event",
			id:      1,
			wantErr: false,
		},
		{
			name:    "non-existing event",
			id:      999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEvent(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("GetEvent() returned nil for existing event")
			}
		})
	}
}
