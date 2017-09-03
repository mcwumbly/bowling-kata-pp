package game_test

import (
	"github.com/mcwumbly/bowl-kata-pp-01/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		app game.Game
	)
	BeforeEach(func() {
		app = game.Game{}
	})

	Describe("AddBowl", func() {
		It("Adds a bowl and returns the list of bowls", func() {
			bowls, err := app.AddBowl([]int{}, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(bowls).To(ConsistOf(3))

			bowls, err = app.AddBowl([]int{5}, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(bowls).To(ConsistOf(5, 3))

			bowls, err = app.AddBowl([]int{3, 4, 10, 7}, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(bowls).To(ConsistOf(3, 4, 10, 7, 3))
		})

		Context("When the number of pins is less than 0", func() {
			It("returns an error", func() {
				_, err := app.AddBowl([]int{}, -1)
				Expect(err).To(MatchError("number of pins cannot be less than 0"))
			})
		})

		Context("When the number of pins is greater than 10", func() {
			It("returns an error", func() {
				_, err := app.AddBowl([]int{}, 11)
				Expect(err).To(MatchError("number of pins cannot be greater than 10"))
			})
		})

		Context("When the number of pins is greater than the number of pins left in the frame", func() {
			It("returns an error", func() {
				_, err := app.AddBowl([]int{3}, 8)
				Expect(err).To(MatchError("number of pins cannot be greater than remaining pins: 7"))
			})

			It("returns an error", func() {
				_, err := app.AddBowl([]int{3, 4, 10, 7}, 4)
				Expect(err).To(MatchError("number of pins cannot be greater than remaining pins: 3"))
			})
		})
	})

	Describe("Frames", func() {
		It("Converts the bowls to frames", func() {
			bowls := []int{}
			frames := app.Frames(bowls)
			Expect(len(frames)).To(Equal(0))

			bowls = []int{3, 4}
			frames = app.Frames(bowls)
			Expect(len(frames)).To(Equal(1))
		})
	})
})
