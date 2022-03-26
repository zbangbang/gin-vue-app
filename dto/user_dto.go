package dto

import "zbangbang/gin-vue-app/model"

type UserDto struct {
	UserName  string `json:"name"`
	Telephone string `json:"telephone"`
}

// 返回部分信息
func ToUserDto(user model.User) UserDto {
	return UserDto{
		UserName:  user.UserName,
		Telephone: user.Telephone,
	}
}
