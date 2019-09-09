package service

import (
	"iris-init/config"
	"github.com/kataras/iris/core/errors"
)

// get all
func GetAll(model interface{}) error {
	err := config.OrmConn.Find(model).Error
	if err != nil {
		return errors.New("select all error")
	}
	return nil
}

// get by id
func GetById(model interface{}, id int64) error {
	if err := config.OrmConn.First(model, id).Error; err != nil {
		return errors.New("select one error")
	}
	return nil
}

// condition select
func GetWithCondition(model interface{}, condition string) error {
	if err := config.OrmConn.Where(condition).Find(model).Error; err != nil {
		return errors.New("no content")
	}
	return nil
}

// update data
func UpdateData(model interface{}, data interface{}) error {
	if err := config.OrmConn.Model(model).Updates(data).Error; err != nil {
		return errors.New("Failed to Update data")
	}
	return nil
}
