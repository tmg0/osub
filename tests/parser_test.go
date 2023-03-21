package tests

import (
	"osub/pkg/parser"
	"testing"
	"time"
)

func TestIntervalParser(t *testing.T) {
	expected := 10 * time.Second

	duration, _ := parser.Interval("10s")

	if duration != expected {
		t.Errorf("Addition test failed: got %d, expected %d", duration, expected)
	}
}
