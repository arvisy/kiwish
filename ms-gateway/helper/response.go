package helper

type Response struct {
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

type Message struct {
	Message string `json:"message"`
}
