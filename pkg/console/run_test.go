package console

import (
	"osub/pkg/shared"
	"testing"
)

func TestRun(t *testing.T) {
	conf, _ := shared.ReadConfig()
	Setup(conf)
}
