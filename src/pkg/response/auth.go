package response

import "github.com/google/uuid"

type UserInfo struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
}

type SignInReponse struct {
	AccessToken string   `json:"access_token"`
	User        UserInfo `json:"user"`
}
