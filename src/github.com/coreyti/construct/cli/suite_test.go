package cli_test

import (
	"bytes"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cli")
}

var originalArgs []string

func clearArgs() {
	originalArgs = os.Args[:]
	os.Args = os.Args[:1]
}

func resetArgs() {
	os.Args = originalArgs[:]
}

func executeCommand(command *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(command, args...)
	return output, err
}

func executeCommandC(command *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	command.SetOutput(buf)
	command.SetArgs(args)

	c, err = command.ExecuteC()

	return c, buf.String(), err
}
