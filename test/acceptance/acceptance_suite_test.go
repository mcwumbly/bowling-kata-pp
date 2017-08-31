package acceptance_test

import (
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
	session *gexec.Session
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
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("bowling kata++", func() {
	BeforeEach(func() {
		var err error
		cmd := exec.Command(binPath)
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		session.Interrupt()
		Eventually(session, TIMEOUT).Should(gexec.Exit(0))
		Eventually(session.Out).Should(gbytes.Say("exiting..."))
	})

	It("prints to stdout when it starts up", func() {
		Eventually(session.Out).Should(gbytes.Say("let's start bowling!"))
	})
})
