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
	if len(bowls) > 0 {
		total := 0
		balls := []view.Ball{}
		for i, bowl := range bowls {
			balls = append(balls, view.Ball{
				Ball: len(balls) + 1,
				Pins: bowl,
			})
			total += bowl
			currentFrame := g.CurrentFrame(bowls[:i])
			if len(balls) == 1 {
				frames = append(frames, view.Frame{
					Frame: currentFrame,
					Balls: balls,
					Total: total,
				})
			} else {
				frames[currentFrame-1].Balls = balls
				frames[currentFrame-1].Total = total
			}
			if len(balls) != 1 || total == 10 {
				total = 0
				balls = []view.Ball{}
			}
		}
	}
	return frames
}

func (g Game) CurrentFrame(bowls []int) int {
	currentFrame := 1
	ball, remaining := 1, 10
	for _, bowl := range bowls {
		remaining -= bowl
		if remaining == 0 || ball == 2 {
			ball, remaining = 1, 10
			currentFrame++
		} else {
			ball = 2
		}
	}
	if currentFrame > 10 {
		return 0
	}
	return currentFrame
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
