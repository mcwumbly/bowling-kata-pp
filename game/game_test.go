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

	Describe("CurrentFrame", func() {
		It("Returns the current frame", func() {
			currentFrame := app.CurrentFrame()
			Expect(currentFrame).To(Equal(1))

			app.AddBowl(3)
			currentFrame = app.CurrentFrame()
			Expect(currentFrame).To(Equal(1))

			app.AddBowl(4)
			currentFrame = app.CurrentFrame()
			Expect(currentFrame).To(Equal(2))

			app.AddBowl(10)
			currentFrame = app.CurrentFrame()
			Expect(currentFrame).To(Equal(3))
		})

		Context("when the game is done", func() {
			It("returns 0", func() {
				app = game.New(10, 10, 10, 10, 10, 10, 10, 10, 10, 10)
				currentFrame := app.CurrentFrame()
				Expect(currentFrame).To(Equal(0))
			})
		})
	})

	Describe("RemainingPins", func() {
		It("Returns the remaining pins for the current frame", func() {
			remainingPins := app.RemainingPins()
			Expect(remainingPins).To(Equal(10))

			app.AddBowl(3)
			remainingPins = app.RemainingPins()
			Expect(remainingPins).To(Equal(7))

			app.AddBowl(4)
			remainingPins = app.RemainingPins()
			Expect(remainingPins).To(Equal(10))

			app.AddBowl(10)
			remainingPins = app.RemainingPins()
			Expect(remainingPins).To(Equal(10))
		})

		Context("when the game is done", func() {
			It("returns 0", func() {
				app = game.New(10, 10, 10, 10, 10, 10, 10, 10, 10, 10)
				remainingPins := app.RemainingPins()
				Expect(remainingPins).To(Equal(0))
			})
		})
	})
})
