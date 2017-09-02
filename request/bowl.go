package request

type BowlRequest struct {
	Bowl struct {
		Pins int `json:"pins"`
	} `json:"bowl"`
}
