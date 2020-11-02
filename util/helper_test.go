package util

import (
	"context"
	"testing"
	"time"
)

func TestGetURLResponse(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	resp, _ := GetURLResponse(ctx, "sports")
	if len(resp) == 0 {
		t.Errorf("URL response is not fetched")
	}
}
