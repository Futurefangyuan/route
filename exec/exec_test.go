package exec_test

import (
	"fmt"
	"testing"

	"github.com/weiguoyu/route/exec"
)



func TestExecCmd(t *testing.T) {
	// test read file
	fmt.Println("------ start -----")
	exec.ExecCmd("go", "env")
}

func TestExecCmd2(t *testing.T) {
	// test read file
	fmt.Println("------ start -----")
	exec.ExecCmd("route", "add 121.4.9.76 MASK 255.255.255.255 10.88.88.2  IF 5")
}

