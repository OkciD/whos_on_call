package api

import appModels "github.com/OkciD/whos_on_call/internal/app/models"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
