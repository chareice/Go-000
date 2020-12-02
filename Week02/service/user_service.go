package service

import (
	"user_service/dao"

	"github.com/pkg/errors"
)

// UserInfo 用户信息
type UserInfo struct {
	Name string
	ID   int
}

// GetUserInfo 通过 id 获取用户信息
func GetUserInfo(id int) (*UserInfo, error) {
	userRecord, err := dao.FetchUserByID(id)

	if err != nil {
		return nil, errors.Wrap(err, "user fetch error")
	}

	return &UserInfo{
		Name: userRecord.Name,
		ID:   userRecord.ID,
	}, nil
}
