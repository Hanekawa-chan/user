package httpserver

import (
	"context"
	"encoding/json"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
	"net/http"
)

func (a *adapter) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := models.CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := a.service.CreateUser(ctx, &req)
	if err != nil {
		return
	}

	_, err = w.Write([]byte(id.String()))
	if err != nil {
		return
	}
}
