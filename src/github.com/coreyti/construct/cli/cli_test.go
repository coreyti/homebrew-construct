package cli_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/coreyti/construct/cli"
)

var _ = Describe("construct (root command)", func() {
	BeforeEach(func() {
		clearArgs()
	})

	AfterEach(func() {
		resetArgs()
	})

	It("executes", func() {
		// var commandArgs []string
		command := cli.NewCLI()

		output, err := executeCommand(command)

		Expect(err).To(BeNil())
		Expect(output).To(HavePrefix("A CLI with utilities for MacOS workstations.\n\n"))
	})
})
