package routing

import (
	"encoding/json"
	"net/http"
)

type JsonHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func Json(inner JsonHandlerFunc) FailableHandlerFunc {
	return FailableHandlerFunc(func(writer http.ResponseWriter, request *http.Request) error {
		response, err := inner(writer, request)
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

func JsonHttp(inner JsonHandlerFunc) http.HandlerFunc {
	return Failable(Json(inner))
}