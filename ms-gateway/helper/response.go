package helper

type Response struct {
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}
