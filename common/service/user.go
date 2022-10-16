package service

import "github.com/luoren18/go-dubbo/service_common/model"

type UserService interface {
	GetAllUser() []model.UserResp
	GetUserById(Id int) model.UserResp
}
