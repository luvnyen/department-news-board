package _user

import "github.com/luvnyen/department-news-board/pkg/models"

type UserResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	NRP   string `json:"nrp"`
	Email string `json:"email"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
