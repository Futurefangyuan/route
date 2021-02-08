package read_test

import (
	"fmt"
	"testing"

	"github.com/weiguoyu/route/read"
)



func TestReadFile(t *testing.T) {
	// test read file
	fmt.Println("------ start -----")
	ip_map, _ := read.ReadFile("C:/Users/yuweiguo/Documents/WeChat Files/yu945019856/FileStorage/File/2021-02/IPAddress.ini")
	fmt.Printf("ip_map: %v\n", ip_map)
}
