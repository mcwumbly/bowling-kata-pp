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
	remaining := remainingPins(bowls)
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
				Ball: i + 1,
				Pins: bowl,
			})
			total += bowl
		}
		frames = []view.Frame{{
			Frame: 1,
			Balls: balls,
			Total: total,
		}}
	}
	return frames
}

func remainingPins(bowls []int) int {
	ball, remaining := 1, 10
	for _, bowl := range bowls {
		remaining -= bowl
		if remaining == 0 || ball == 2 {
			ball, remaining = 1, 10
		} else {
			ball = 2
		}
	}
	return remaining
}
