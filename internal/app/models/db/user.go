package db

import appModels "github.com/OkciD/whos_on_call/internal/app/models"

type User struct {
	ID         int
	Name       string
	ApiKeyHash string
}

func (u *User) ToAppModel() *appModels.User {
	return &appModels.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
