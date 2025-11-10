package request

type Conditions struct {
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"0"`
}
