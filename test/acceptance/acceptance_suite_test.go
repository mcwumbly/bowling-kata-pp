package acceptance_test

import (
	"bytes"
	"fmt"
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
		By("getting the initial state of the game", func() {
			resp, err := http.Get(url)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(respBytes)).To(MatchJSON(`{
			"game": {
				"currentFrame": 1,
				"remainingPins": 10,
				"frames": [],
				"total": 0
			}
		}`))
		})

		By("bowling the first ball", func() {
			resp := addBowl(3)
			Expect(resp).To(MatchJSON(`{
			"game": {
				"currentFrame": 1,
				"remainingPins": 7,
				"frames": [
				  {
						"frame": 1,
						"balls": [
						  { "ball": 1, "pins": 3 }
						],
						"total": 3
					}
				],
				"total": 3
			}
		}`))
		})

		By("bowling the second ball", func() {
			resp := addBowl(5)
			Expect(resp).To(MatchJSON(`{
			"game": {
				"currentFrame": 2,
				"remainingPins": 10,
				"frames": [
				  {
						"frame": 1,
						"balls": [
						  { "ball": 1, "pins": 3 },
						  { "ball": 2, "pins": 5 }
						],
						"total": 8
					}
				],
				"total": 8
			}
		}`))
		})

		By("bowling the remaining balls", func() {
			addBowl(4)
			addBowl(6)
			addBowl(1)
			addBowl(2)
			addBowl(5)
			addBowl(5)
			addBowl(10)
			addBowl(10)
			addBowl(10)
			addBowl(10)
			addBowl(10)
			resp := addBowl(10)
			Expect(resp).To(MatchJSON(`{
			"game": {
				"frames": [
				  { "frame": 1, "total": 8,
						"balls": [ { "ball": 1, "pins": 3 }, { "ball": 2, "pins": 5 } ] },
				  { "frame": 2, "total": 10,
						"balls": [ { "ball": 1, "pins": 4 }, { "ball": 2, "pins": 6 } ] },
				  { "frame": 3, "total": 3,
						"balls": [ { "ball": 1, "pins": 1 }, { "ball": 2, "pins": 2 } ] },
				  { "frame": 4, "total": 10,
						"balls": [ { "ball": 1, "pins": 5 }, { "ball": 2, "pins": 5 } ] },
				  { "frame": 5, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] },
				  { "frame": 6, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] },
				  { "frame": 7, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] },
				  { "frame": 8, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] },
				  { "frame": 9, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] },
				  { "frame": 10, "total": 10,
						"balls": [ { "ball": 1, "pins": 10 } ] }
				],
				"total": 91
			}
		}`))
		})
	})

	Context("when you try to bowl more than 10 pins", func() {
		It("returns an error", func() {
			req := bytes.NewBufferString(`{ "bowl": { "pins": 11 } }`)
			resp, err := http.Post(url, "application/json", req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(respBytes)).To(MatchJSON(`{
			"error": "number of pins cannot be greater than 10"
		}`))
		})
	})

	Context("when you try to bowl more than the number of remaining pins", func() {
		It("returns an error", func() {
			req := bytes.NewBufferString(`{ "bowl": { "pins": 8 } }`)
			_, err := http.Post(url, "application/json", req)
			Expect(err).NotTo(HaveOccurred())

			req = bytes.NewBufferString(`{ "bowl": { "pins": 3 } }`)
			resp, err := http.Post(url, "application/json", req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(respBytes)).To(MatchJSON(`{
				"error": "number of pins cannot be greater than remaining pins: 2"
		}`))
		})
	})

	Context("when you the request is invalid json", func() {
		It("returns an error", func() {
			req := bytes.NewBufferString(`{ "bowl": { "pins": 8 } `)
			resp, err := http.Post(url, "application/json", req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(respBytes)).To(MatchJSON(`{
				"error": "unexpected end of JSON input"
		}`))
		})
	})
})

func addBowl(pins int) string {
	req := bytes.NewBufferString(fmt.Sprintf(`{ "bowl": { "pins": %d } }`, pins))
	resp, err := http.Post(url, "application/json", req)
	Expect(err).NotTo(HaveOccurred())
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())
	return string(respBytes)
}
