package view

type Response struct {
	Game Game `json:"game"`
}

type Game struct {
	Frames []Frame `json:"frames"`
	Total  int     `json:"total"`
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
