package server

// helper enabling error message to be returned in JSON format
type jsonError struct {
	Msg string `json:"message"`
}

func (e *jsonError) Error() string {
	return e.Msg
}
