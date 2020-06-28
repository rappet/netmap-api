package routing

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err HttpError) Error() string {
	return err.Message
}
