package game

import (
	"errors"
	"fmt"

	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

type Game struct {
	bowls []int
}

type frame struct {
	number int
	total  int
	balls  []int
}

type game struct {
	bowls []int
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
	m := game{bowls: g.Bowls()}
	for _, f := range m.frames() {
		if len(f.balls) == 0 {
			break
		}
		frames = append(frames, viewFrame(f))
	}
	return frames
}

func viewFrame(f frame) view.Frame {
	balls := []view.Ball{}
	for i, b := range f.balls {
		balls = append(balls, view.Ball{
			Ball: i + 1,
			Pins: b,
		})
	}
	return view.Frame{
		Frame: f.number,
		Total: f.total,
		Balls: balls,
	}
}

func (m game) frames() []frame {
	frames := []frame{frame{number: 1, total: 0, balls: []int{}}}
	f := &frames[len(frames)-1]
	for _, bowl := range m.bowls {
		if f.isComplete() {
			frames = append(frames, frame{number: len(frames) + 1, total: 0, balls: []int{}})
			f = &frames[len(frames)-1]
		}
		f.total += bowl
		f.balls = append(f.balls, bowl)
	}
	return frames
}

func (f frame) isComplete() bool {
	return len(f.balls) == 2 || f.total == 10
}

func (m game) isComplete() bool {
	frames := m.frames()
	return len(frames) == 10 && frames[len(frames)-1].isComplete()
}

func (m game) currentFrame() frame {
	if m.isComplete() {
		return frame{number: 0, total: 10}
	}
	frames := m.frames()
	f := frames[len(frames)-1]
	if f.isComplete() {
		f = frame{number: f.number + 1, balls: []int{}}
	}
	return f
}

func (g *Game) CurrentFrame() int {
	m := game{bowls: g.Bowls()}
	return m.currentFrame().number
}

func (g *Game) RemainingPins() int {
	m := game{bowls: g.Bowls()}
	return 10 - m.currentFrame().total
}
