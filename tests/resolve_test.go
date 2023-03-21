package tests

import (
	"osub/pkg/resolve"
	"testing"
	"time"
)

func TestResolveInterval(t *testing.T) {
	expected := 10 * time.Second

	duration, _ := resolve.Interval("10s")

	if duration != expected {
		t.Errorf("Resolve interval test failed: got %d, expected %d", duration, expected)
	}
}
