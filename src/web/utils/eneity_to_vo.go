package utils

import (
	"fmt"
	"reflect"
	"time"
)

// EntityToVO
//
//	@Description: 将实体转换为VO(支持切片和非切片的转换,时间也会处理成字符串)
//	@param entity 实体或切片
//	@param vo VO
//	@return error 错误
func EntityToVO(entity any, vo any) error {
	entityType := reflect.TypeOf(entity)
	entityValue := reflect.ValueOf(entity)
	voType := reflect.TypeOf(vo)
	voValue := reflect.ValueOf(vo)

	// 确保传入的是指针类型
	if entityType.Kind() != reflect.Ptr || voType.Kind() != reflect.Ptr {
		return fmt.Errorf("both entity and vo must be pointers")
	}

	// 获取实际值
	entityValue = entityValue.Elem()
	voValue = voValue.Elem()

	// 判断是否是切片转换
	if entityValue.Kind() == reflect.Slice && voValue.Kind() == reflect.Slice {
		if voValue.IsNil() {
			voValue.Set(reflect.MakeSlice(voValue.Type(), 0, entityValue.Len()))
		}

		// 遍历实体切片并转换
		for i := 0; i < entityValue.Len(); i++ {
			entityItem := entityValue.Index(i).Interface()
			voItem := reflect.New(voValue.Type().Elem().Elem()).Interface()

			// 转换单个结构体
			err := EntityToVO(entityItem, voItem)
			if err != nil {
				return err
			}

			// 追加指针类型的值
			voValue.Set(reflect.Append(voValue, reflect.ValueOf(voItem)))
		}
	} else if entityValue.Kind() == reflect.Struct && voValue.Kind() == reflect.Struct {
		// 单个结构体转换
		for i := 0; i < entityValue.NumField(); i++ {
			field := entityValue.Type().Field(i)
			entityField := entityValue.Field(i)

			// 查找 VO 中是否有同名字段
			voField := voValue.FieldByName(field.Name)
			if voField.IsValid() && voField.CanSet() {
				// 如果类型相同，则直接赋值
				if voField.Type() == entityField.Type() {
					voField.Set(entityField)
				}
			}
		}

		// 特殊处理时间字段
		handleTimeField(entityValue, voValue)
	} else {
		return fmt.Errorf("both entity and vo must be either structs or slices of structs")
	}

	return nil
}
func handleTimeField(entityValue reflect.Value, voValue reflect.Value) {
	timeFormat := "2006-05-04 15:02:01"

	// 处理 CreatedAt -> CreatedTime
	if createdAtField := entityValue.FieldByName("CreatedAt"); createdAtField.IsValid() && createdAtField.Type() == reflect.TypeOf(time.Time{}) {
		if createdTimeField := voValue.FieldByName("CreatedTime"); createdTimeField.IsValid() && createdTimeField.CanSet() {
			createdTimeField.SetString(createdAtField.Interface().(time.Time).Format(timeFormat))
		}
	}

	// 处理 UpdatedAt -> UpdatedTime
	if updatedAtField := entityValue.FieldByName("UpdatedAt"); updatedAtField.IsValid() {
		if updatedAtField.Type() == reflect.TypeOf(&time.Time{}) {
			// 如果 UpdatedAt 是指针类型，则需要进行空指针检查
			if updatedAtField.Interface() != nil {
				// 转换并赋值
				if updatedTimeField := voValue.FieldByName("UpdatedTime"); updatedTimeField.IsValid() && updatedTimeField.CanSet() {
					// 使用类型断言并检查 nil
					if updatedAtPtr, ok := updatedAtField.Interface().(*time.Time); ok && updatedAtPtr != nil {
						updatedTimeField.SetString(updatedAtPtr.Format(timeFormat))
					} else {
						updatedTimeField.SetString("") // 如果指针为 nil，设置为空字符串
					}
				}
			} else {
				// 如果 UpdatedAt 为 nil，设置空字符串
				if updatedTimeField := voValue.FieldByName("UpdatedTime"); updatedTimeField.IsValid() && updatedTimeField.CanSet() {
					updatedTimeField.SetString("")
				}
			}
		}
	}
}
