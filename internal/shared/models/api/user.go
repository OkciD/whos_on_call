package api

import appModels "github.com/OkciD/whos_on_call/internal/shared/models"

func (u *User) ToAppModel() *appModels.User {
	return &appModels.User{
		ID:   int(u.Id),
		Name: u.Name,
	}
}

func FromUserAppModel(appUser *appModels.User) *User {
	return &User{
		//nolint:gosec // todo:fix potential overflow
		Id:   int32(appUser.ID),
		Name: appUser.Name,
	}
}
