package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	UserService  = userService{}
)

type userService struct {
}

func (service userService) Create(user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.Create(user)
	 return err1
}

func (service userService) Login(auser *model.LoginUser) (*model.Auth, *httperrors.HttpError) {
	user, err1 :=  r.Userrepository.Login(auser)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}
func (service userService) Logout(token string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	success, failure := r.Userrepository.Logout(token)
	return success, failure
}
func (service userService) GetOne(code string) (*model.User, *httperrors.HttpError) {
	user, err1 := r.Userrepository.GetOne(code)
	return user, err1
}

func (service userService) GetAll(search string) ([]*model.User, *httperrors.HttpError) {
	users, err := r.Userrepository.GetAll(search)
	return users, err
}

func (service userService) Update(code string, user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.Update(code, user)
	return err1
}

func (service userService) AUpdate(id string, user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.AUpdate(id, user)
	return err1
}
func (service userService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Userrepository.Delete(id)
		return success, failure
}
