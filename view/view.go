package view

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
