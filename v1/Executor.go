package v1

import (
	"bytes"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var KeyPathFlag string
var Json bool
var Raw bool
var JsonPath string

func ExecuteCmd(root *cobra.Command, args ...string) (stdout string, stderr string, err error) {
	var stderrBuf, stdoutBuf bytes.Buffer
	root.SetOutput(&stderrBuf)
	if len(args) == 0 {
		args = make([]string, 0)
	}
	root.SetArgs(args)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	_, err = root.ExecuteC()

	_ = w.Close()
	os.Stdout = old
	_, _ = io.Copy(&stdoutBuf, r)

	ResetSharedFlags()

	return stdoutBuf.String(), stderrBuf.String(), err
}


func ResetSharedFlags() {
	KeyPathFlag = ""
	Json = false
	Raw = false
	JsonPath = ""
}