package game_test

import (
	"github.com/mcwumbly/bowl-kata-pp-01/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		app *game.Game
	)
	BeforeEach(func() {
		app = game.New()
	})

	Describe("AddBowl", func() {
		It("Adds a bowl and returns the list of bowls", func() {
			err := app.AddBowl(3)
			Expect(err).NotTo(HaveOccurred())
			Expect(app.Bowls()).To(ConsistOf(3))

			err = app.AddBowl(5)
			Expect(err).NotTo(HaveOccurred())
			Expect(app.Bowls()).To(ConsistOf(3, 5))

			err = app.AddBowl(10)
			Expect(err).NotTo(HaveOccurred())
			Expect(app.Bowls()).To(ConsistOf(3, 5, 10))

			err = app.AddBowl(7)
			Expect(err).NotTo(HaveOccurred())
			Expect(app.Bowls()).To(ConsistOf(3, 5, 10, 7))
		})

		Context("When the number of pins is less than 0", func() {
			It("returns an error", func() {
				err := app.AddBowl(-1)
				Expect(err).To(MatchError("number of pins cannot be less than 0"))
			})
		})

		Context("When the number of pins is greater than 10", func() {
			It("returns an error", func() {
				err := app.AddBowl(11)
				Expect(err).To(MatchError("number of pins cannot be greater than 10"))
			})
		})

		Context("When the number of pins is greater than the number of pins left in the frame", func() {
			It("returns an error", func() {
				err := app.AddBowl(3)
				Expect(err).NotTo(HaveOccurred())
				err = app.AddBowl(8)
				Expect(err).To(MatchError("number of pins cannot be greater than remaining pins: 7"))
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
			Expect(frames[0].Frame).To(Equal(1))
			Expect(frames[0].Total).To(Equal(7))
			Expect(len(frames[0].Balls)).To(Equal(2))
			Expect(frames[0].Balls[0].Ball).To(Equal(1))
			Expect(frames[0].Balls[0].Pins).To(Equal(3))
			Expect(frames[0].Balls[1].Ball).To(Equal(2))
			Expect(frames[0].Balls[1].Pins).To(Equal(4))

			bowls = []int{3, 4, 5}
			frames = app.Frames(bowls)
			Expect(len(frames)).To(Equal(2))
			Expect(frames[1].Frame).To(Equal(2))
			Expect(frames[1].Total).To(Equal(5))
			Expect(len(frames[1].Balls)).To(Equal(1))
			Expect(frames[1].Balls[0].Ball).To(Equal(1))
			Expect(frames[1].Balls[0].Pins).To(Equal(5))

			bowls = []int{3, 4, 5, 5, 10, 10}
			frames = app.Frames(bowls)
			Expect(len(frames)).To(Equal(4))
			Expect(frames[3].Frame).To(Equal(4))
			Expect(frames[3].Total).To(Equal(10))
			Expect(len(frames[3].Balls)).To(Equal(1))
			Expect(frames[3].Balls[0].Ball).To(Equal(1))
			Expect(frames[3].Balls[0].Pins).To(Equal(10))
		})
	})

	Describe("CurrentFrame", func() {
		It("Returns the current frame", func() {
			bowls := []int{}
			currentFrame := app.CurrentFrame(bowls)
			Expect(currentFrame).To(Equal(1))

			bowls = []int{3}
			currentFrame = app.CurrentFrame(bowls)
			Expect(currentFrame).To(Equal(1))

			bowls = []int{3, 4}
			currentFrame = app.CurrentFrame(bowls)
			Expect(currentFrame).To(Equal(2))

			bowls = []int{3, 4, 10}
			currentFrame = app.CurrentFrame(bowls)
			Expect(currentFrame).To(Equal(3))
		})

		Context("when the game is done", func() {
			It("returns 0", func() {
				bowls := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
				currentFrame := app.CurrentFrame(bowls)
				Expect(currentFrame).To(Equal(0))
			})
		})
	})

	Describe("RemainingPins", func() {
		It("Returns the remaining pins for the current frame", func() {
			bowls := []int{}
			remainingPins := app.RemainingPins(bowls)
			Expect(remainingPins).To(Equal(10))

			bowls = []int{3}
			remainingPins = app.RemainingPins(bowls)
			Expect(remainingPins).To(Equal(7))

			bowls = []int{3, 4}
			remainingPins = app.RemainingPins(bowls)
			Expect(remainingPins).To(Equal(10))

			bowls = []int{3, 4, 10}
			remainingPins = app.RemainingPins(bowls)
			Expect(remainingPins).To(Equal(10))
		})

		Context("when the game is done", func() {
			It("returns 0", func() {
				bowls := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
				remainingPins := app.RemainingPins(bowls)
				Expect(remainingPins).To(Equal(0))
			})
		})
	})
})
