package httpserver

import (
	"context"
	"encoding/json"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
	"net/http"
)

func (a *adapter) createUser(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	req := models.CreateUserRequest{}
	resp := models.CreateUserResponse{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	id, err := a.service.CreateUser(ctx, &req)
	if err != nil {
		return err
	}

	resp.UserId = id.String()

	err = sendResponse(w, resp)

	return err
}
