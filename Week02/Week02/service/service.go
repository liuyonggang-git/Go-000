package service

import (
	"Week02/dao"
	"database/sql"
	"errors"
)

// 根据业务需要，直接抛
func GetUsernameById1(id string) (*string, error) {
	return dao.GetUsernameById(id)
}

// 根据业务需要，查不到时返回nil
func GetUsernameById2(id string) (*string, error) {
	username, err := dao.GetUsernameById(id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return username, nil
	} else {
		return username, err
	}
}
