package routing

import (
	"encoding/json"
	"net/http"
)

type JsonHandlerFunc func(*http.Request) (interface{}, error)

func Json(inner JsonHandlerFunc) FailableHandlerFunc {
	return FailableHandlerFunc(func(writer http.ResponseWriter, request *http.Request) error {
		response, err := inner(request)
		if err != nil {
			return err
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			return err
		}

		return nil
	})
}