package acceptance_test

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	binPath string
)

const TIMEOUT = "5s"

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var _ = BeforeSuite(func() {
	var err error
	srcPath := os.Getenv("GOPATH") + "/src/github.com/mcwumbly/bowl-kata-pp-01/main.go"
	binPath, err = gexec.Build(srcPath, "-race")
	fmt.Printf("wrote binary to %s\n", binPath)
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("bowling kata++", func() {
	It("prints to stdout when it starts up", func() {
		cmd := exec.Command(binPath)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, TIMEOUT).Should(gexec.Exit(0))
		Expect(session.Out).To(gbytes.Say("let's start bowling!"))
	})
})
