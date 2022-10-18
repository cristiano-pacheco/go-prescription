package controller

import (
	"github.com/cristiano-pacheco/go-prescription/model"
)

type UserActions struct {
	CreateAction  *userCreateAction
	DestroyAction *userDestroyAction
	EditAction    *userEditAction
	IndexAction   *userIndexAction
	StoreAction   *userStoreAction
	UpdateAction  *userUpdateAction
}

func CreateUserActions(userModel *model.UserModel) *UserActions {
	return &UserActions{
		CreateAction:  NewUserCreateAction(userModel),
		DestroyAction: NewUserDestroyAction(userModel),
		EditAction:    NewUserEditAction(userModel),
		IndexAction:   NewUserIndexAction(userModel),
		StoreAction:   NewUserStoreAction(userModel),
		UpdateAction:  NewUserUpdateAction(userModel),
	}
}
