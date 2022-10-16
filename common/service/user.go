package service

import "luoren.top/go_dubbo/service_common/model"

type UserService interface {
	GetAllUser() []model.UserResp
	GetUserById(Id int) model.UserResp
}
