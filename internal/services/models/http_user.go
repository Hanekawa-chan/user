package models

type CreateUserRequest struct {
	Email   string `json:"email"`
	Country string `json:"country"`
}

type CreateUserResponse struct {
	UserId string `json:"user_id"`
}
