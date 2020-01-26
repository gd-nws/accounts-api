package Models

type ErrorResponse struct {
	Message string `json:"message"`
	Trace   string `json:"trace"`
}
