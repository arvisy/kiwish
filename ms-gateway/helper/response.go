package helper

type Response struct {
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}
