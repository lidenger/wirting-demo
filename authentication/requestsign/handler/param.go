package handler

type GenAccessTokenParam struct {
	Sign string `json:"sign"`
	Ak   string `json:"ak"`
	Sk   string `json:"sk"`
}
