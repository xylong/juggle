package validator

import (
	"github.com/go-playground/validator/v10"
	"juggle/basic/model"
	"log"
	"time"
)

func init() {
	if err := v.RegisterValidation("birthday", birthday); err != nil {
		log.Fatalln("validator error:" + err.Error())
	}
	if err := v.RegisterValidation("size", size); err != nil {
		log.Fatalln("validator error:" + err.Error())
	}
}

var (
	// birthday 出生日期验证规则
	birthday validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(string)

		if !ok || len(value) != 10 {
			return false
		}

		t, err := time.Parse("2006-01-02", value)
		if err != nil {
			return false
		}

		age := time.Now().Year() - t.Year()
		if age < 0 || age > 100 {
			return false
		}

		return true
	}

	// size 验证切片数量
	size validator.Func = func(fl validator.FieldLevel) bool {
		users, ok := fl.Parent().Interface().(model.Users)
		if !ok {
			return false
		}

		list, ok := fl.Field().Interface().([]model.User)
		if !ok {
			return false
		}

		return users.Size > 0 && users.Size == len(list)
	}
)
