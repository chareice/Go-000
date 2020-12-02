package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// User 用户数据结构
type User struct {
	ID   int
	Name string
}

// FetchUserByID 通过 id 查找 User 记录
func FetchUserByID(id int) (*User, error) {
	if id > 0 {
		user := &User{}
		user.ID = id
		user.Name = "test"
		return user, nil
	}

	return nil, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("users %d record not found", id))
}
