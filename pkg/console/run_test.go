package console

import (
	"fmt"
	"os/exec"
	"osub/pkg/shared"
	"testing"
)

func TestRun(t *testing.T) {
	exec.Command("pwd").Run()
	conf, _ := shared.ReadConfig("../../")
	servers, _ := setup(conf)
	fmt.Println(servers)
}
