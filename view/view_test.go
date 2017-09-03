package view_test

import (
	"github.com/mcwumbly/bowl-kata-pp-01/game"
	"github.com/mcwumbly/bowl-kata-pp-01/view"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("View", func() {
	var (
		app *game.Game
	)
	BeforeEach(func() {
		app = game.New()
	})

	Describe("Frames", func() {
		It("Converts the frames to view frames", func() {
			frames := view.Frames(app.Frames())
			Expect(len(frames)).To(Equal(0))

			app = game.New(3, 4)
			frames = view.Frames(app.Frames())
			Expect(len(frames)).To(Equal(1))
			Expect(frames[0].Frame).To(Equal(1))
			Expect(frames[0].Total).To(Equal(7))
			Expect(len(frames[0].Balls)).To(Equal(2))
			Expect(frames[0].Balls[0].Ball).To(Equal(1))
			Expect(frames[0].Balls[0].Pins).To(Equal(3))
			Expect(frames[0].Balls[1].Ball).To(Equal(2))
			Expect(frames[0].Balls[1].Pins).To(Equal(4))

			app = game.New(3, 4, 5)
			frames = view.Frames(app.Frames())
			Expect(len(frames)).To(Equal(2))
			Expect(frames[1].Frame).To(Equal(2))
			Expect(frames[1].Total).To(Equal(5))
			Expect(len(frames[1].Balls)).To(Equal(1))
			Expect(frames[1].Balls[0].Ball).To(Equal(1))
			Expect(frames[1].Balls[0].Pins).To(Equal(5))

			app = game.New(3, 4, 5, 5, 10, 10)
			frames = view.Frames(app.Frames())
			Expect(len(frames)).To(Equal(4))
			Expect(frames[3].Frame).To(Equal(4))
			Expect(frames[3].Total).To(Equal(10))
			Expect(len(frames[3].Balls)).To(Equal(1))
			Expect(frames[3].Balls[0].Ball).To(Equal(1))
			Expect(frames[3].Balls[0].Pins).To(Equal(10))
		})
	})
})
