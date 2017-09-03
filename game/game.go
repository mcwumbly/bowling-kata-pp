package game

import (
	"errors"
	"fmt"

	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

type Game struct{}

func (g Game) AddBowl(bowls []int, pins int) ([]int, error) {
	if pins < 0 {
		return bowls, errors.New("number of pins cannot be less than 0")
	}
	if pins > 10 {
		return bowls, errors.New("number of pins cannot be greater than 10")
	}
	remaining := g.RemainingPins(bowls)
	if pins > remaining {
		return bowls, fmt.Errorf("number of pins cannot be greater than remaining pins: %d", remaining)
	}
	return append(bowls, pins), nil
}

func (g Game) Frames(bowls []int) []view.Frame {
	frames := []view.Frame{}
	if len(bowls) == 0 {
		return frames
	}
	frames = append(frames, view.Frame{Balls: []view.Ball{}})
	frame := &frames[len(frames)-1]
	for _, bowl := range bowls {
		if len(frame.Balls) == 2 || frame.Total == 10 {
			frames = append(frames, view.Frame{Balls: []view.Ball{}})
			frame = &frames[len(frames)-1]
		}
		frame.Frame = len(frames)
		frame.Total += bowl
		frame.Balls = append(frame.Balls, view.Ball{
			Ball: len(frame.Balls) + 1,
			Pins: bowl,
		})
	}
	return frames
}

func (g Game) CurrentFrame(bowls []int) int {
	frames := g.Frames(bowls)
	if len(frames) == 0 {
		return 1
	}
	lastFrame := frames[len(frames)-1]
	if len(lastFrame.Balls) == 2 || lastFrame.Total == 10 {
		currentFrame := len(frames) + 1
		if currentFrame > 10 {
			return 0
		}
		return currentFrame
	}
	return len(frames)
}

func (g Game) RemainingPins(bowls []int) int {
	ball, remaining := 1, 10
	for _, bowl := range bowls {
		remaining -= bowl
		if remaining == 0 || ball == 2 {
			ball, remaining = 1, 10
		} else {
			ball = 2
		}
	}
	if g.CurrentFrame(bowls) == 0 {
		return 0
	}
	return remaining
}
