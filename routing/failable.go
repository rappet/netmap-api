package routing

import (
	"encoding/json"
	"log"
	"net/http"
)

type FailableHandlerFunc func(http.ResponseWriter, *http.Request) error

func Failable(inner FailableHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := inner(writer, request)

		if err != nil {
			var httpErr HttpError

			switch typedErr := err.(type) {
			case HttpError:
				httpErr = typedErr
			default:
				httpErr = HttpError{
					Message: err.Error(),
					Code:    http.StatusInternalServerError,
				}
			}

			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			writer.WriteHeader(httpErr.Code)
			if err = json.NewEncoder(writer).Encode(httpErr); err != nil {
				log.Fatalln(err)
			}
		}
	})
}