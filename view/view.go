package view

import "github.com/mcwumbly/bowl-kata-pp-01/game"

type Response struct {
	Game Game `json:"game"`
}

type Game struct {
	CurrentFrame  int     `json:"currentFrame,omitempty"`
	RemainingPins int     `json:"remainingPins,omitempty"`
	Frames        []Frame `json:"frames"`
	Total         int     `json:"total"`
}

type Frame struct {
	Frame int    `json:"frame"`
	Balls []Ball `json:"balls"`
	Total int    `json:"total"`
}

type Ball struct {
	Ball int `json:"ball"`
	Pins int `json:"pins"`
}

func Frames(fs []game.Frame) []Frame {
	frames := []Frame{}
	for _, f := range fs {
		if len(f.Balls) == 0 {
			break
		}
		frames = append(frames, viewFrame(f))
	}
	return frames
}

func viewFrame(f game.Frame) Frame {
	balls := []Ball{}
	for i, b := range f.Balls {
		balls = append(balls, Ball{
			Ball: i + 1,
			Pins: b,
		})
	}
	return Frame{
		Frame: f.Number,
		Total: f.Total,
		Balls: balls,
	}
}
