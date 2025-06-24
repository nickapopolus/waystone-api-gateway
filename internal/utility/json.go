package utility

import (
	"encoding/json"
	"errors"
	"net/http"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	err = decoder.Decode(&struct{}{})
	if err == nil {
		return errors.New("Body must contain a single JSON object")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for _, header := range headers {
		for key, value := range header {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func ErrJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()
	WriteJSON(w, statusCode, payload)
	return nil
}
