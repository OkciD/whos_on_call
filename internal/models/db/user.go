package db

import appModels "github.com/OkciD/whos_on_call/internal/models"

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

func FromUserAppModel(appUser *appModels.User) *User {
	return &User{
		ID:   appUser.ID,
		Name: appUser.Name,
	}
}
