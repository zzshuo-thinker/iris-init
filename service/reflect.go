package service

import (
	"reflect"
	"fmt"
	)

// 设置结构体中的变量
func setField(obj interface{}, name string, value interface{}) error {
	// 结构信息
	structData := reflect.ValueOf(obj).Elem()
	// 通过名字找到结构体中的变量
	fieldValue := structData.FieldByName(name)

	if !fieldValue.IsValid() {
		fmt.Println("No such field ", name)
		return fmt.Errorf("no such field %s", name)
	}

	if !fieldValue.CanSet() {
		fmt.Println("Can not set ", name)
		return fmt.Errorf("can not set %s", name)
	}

	// 结构体中变量的类型
	fieldType := fieldValue.Type()
	// 参数的值
	val := reflect.ValueOf(value)
	// 参数的类型
	valTypeStr := val.Type().String()
	// 结构体中变量的类型
	fieldTypeStr := fieldType.String()
	// float64 to int
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	}
	// 类型必须匹配
	if fieldType != val.Type() {
		return fmt.Errorf("value type %s didn't match obj field type %s ", valTypeStr, fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}

func SetObjectByJSON(obj interface{}, data map[string]interface{}) error {
	for key, value := range data {
		if err := setField(obj, key, value); err != nil {
			fmt.Println("SetObjectByJSON Set field fail.")
			return err
		}
	}
	return nil
}
