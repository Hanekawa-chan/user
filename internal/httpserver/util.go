package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/Hanekawa-chan/kanji-user/internal/services/errors"
	"net/http"
)

type ErrorResponse struct {
	StatusCode     int    `json:"status_code"`
	Error          string `json:"error"`
	LocalizedError string `json:"localized_error"`
}

func JError(w http.ResponseWriter, err error) error {
	localizedError := "Действие не выполнено по неизвестной причине"
	statusCode := http.StatusInternalServerError

	switch err { //noling:errorlint
	case errors.ErrInternal:
		localizedError = "Действие не выполнено по неизвестной причине"
	case errors.ErrValidation:
		localizedError = "Неправильный синтаксис параметров или аргументов"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(
		ErrorResponse{
			Error:          err.Error(),
			LocalizedError: localizedError,
			StatusCode:     statusCode,
		}); err != nil {
		return fmt.Errorf("cannot write response: %w", err)
	}
	return nil
}
