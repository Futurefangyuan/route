package exec

import (
	"os/exec"
	"github.com/sirupsen/logrus"
)


// ExecCmd ...
func ExecCmd(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("err: %v", err)
		return err
	}
	logrus.Infof("combined out:\n%s\n", string(out))
	return 
}


