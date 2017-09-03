package game

import (
	"errors"
	"fmt"

	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

type Game struct {
	bowls []int
}

type Frame struct {
	Number int
	Total  int
	Balls  []int
}

func New(bowls ...int) *Game {
	g := &Game{}
	for _, bowl := range bowls {
		g.AddBowl(bowl)
	}
	return g
}

func (g *Game) Bowls() []int {
	return g.bowls
}

func (g *Game) AddBowl(pins int) error {
	if pins < 0 {
		return errors.New("number of pins cannot be less than 0")
	}
	if pins > 10 {
		return errors.New("number of pins cannot be greater than 10")
	}
	if pins > g.RemainingPins() {
		return fmt.Errorf("number of pins cannot be greater than remaining pins: %d", g.RemainingPins())
	}
	g.bowls = append(g.bowls, pins)
	return nil
}

func (g *Game) Frames() []view.Frame {
	frames := []view.Frame{}
	for _, f := range g.frames() {
		if len(f.Balls) == 0 {
			break
		}
		frames = append(frames, viewFrame(f))
	}
	return frames
}

func viewFrame(f Frame) view.Frame {
	balls := []view.Ball{}
	for i, b := range f.Balls {
		balls = append(balls, view.Ball{
			Ball: i + 1,
			Pins: b,
		})
	}
	return view.Frame{
		Frame: f.Number,
		Total: f.Total,
		Balls: balls,
	}
}

func (g *Game) frames() []Frame {
	frames := []Frame{Frame{Number: 1, Total: 0, Balls: []int{}}}
	f := &frames[len(frames)-1]
	for _, bowl := range g.Bowls() {
		if f.isComplete() {
			frames = append(frames, Frame{Number: len(frames) + 1, Total: 0, Balls: []int{}})
			f = &frames[len(frames)-1]
		}
		f.Total += bowl
		f.Balls = append(f.Balls, bowl)
	}
	return frames
}

func (f Frame) isComplete() bool {
	return len(f.Balls) == 2 || f.Total == 10
}

func (g *Game) isComplete() bool {
	frames := g.frames()
	return len(frames) == 10 && frames[len(frames)-1].isComplete()
}

func (g *Game) currentFrame() Frame {
	if g.isComplete() {
		return Frame{Number: 0, Total: 10}
	}
	frames := g.frames()
	f := frames[len(frames)-1]
	if f.isComplete() {
		f = Frame{Number: f.Number + 1, Balls: []int{}}
	}
	return f
}

func (g *Game) CurrentFrame() int {
	return g.currentFrame().Number
}

func (g *Game) RemainingPins() int {
	return 10 - g.currentFrame().Total
}
