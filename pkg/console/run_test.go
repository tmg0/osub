package console

import (
	"fmt"
	"osub/pkg/shared"
	"testing"
)

func TestRun(t *testing.T) {
	conf, _ := shared.ReadConfig("../../")
	servers, _ := setup(conf)
	fmt.Println(servers)
}
