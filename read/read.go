package read

import (
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)


// ReadFile read the address
func ReadFile(filename string) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)

	val, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Fatal("read file:", err.Error())
	}
	lines := strings.Split(strings.Replace(string(val), "\r\n", "\n", -1), "\n")
	// logrus.Infof("lines: %v, len(lines): %d", lines, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if (line == "") {
			continue
		}
		ip_list := strings.Split(line, "|")
		if (len(ip_list) != 2){
			logrus.Errorf("illegal line: %s", line)
		}
		m[ip_list[0]] = ip_list[1]
	}
	return m, nil
}

