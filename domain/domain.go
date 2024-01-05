package domain

type ErrResponse struct {
	Error string `json:"error"`
}

type DataResponse struct {
	Data []string `json:"data"`
}
