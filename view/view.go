package view

type Response struct {
	Game Game `json:"game"`
}

type Game struct {
	Frames []Frame `json:"frames"`
	Total  int     `json:"total"`
}

type Frame struct {
}
