package acceptance_test

import (
	"io/ioutil"
	"net/http"
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
	url     string
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
	url = "http://localhost:8080/"
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("bowling kata++", func() {
	BeforeEach(func() {
		var err error
		cmd := exec.Command(binPath)
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session.Out, TIMEOUT).Should(gbytes.Say("let's start bowling!"))
	})

	AfterEach(func() {
		session.Interrupt()
		Eventually(session.Out).Should(gbytes.Say("shutting down..."))
		Eventually(session, TIMEOUT).Should(gexec.Exit(0))
		Eventually(session.Out).Should(gbytes.Say("exiting..."))
	})

	It("displays the total score and the completed frames", func() {
		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())
		respBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(respBytes)).To(MatchJSON(`{
			"game": {
				"frames": [],
				"total": 0
			}
		}`))
	})
})
