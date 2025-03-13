package initialize

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var myValidator = myValidate{}

type myValidate struct {
	Validator *validator.Validate
}

func Validator() myValidate {
	return myValidator
}

func Validate() {
	validatorObj := validator.New()
	validatorObj.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("msg")
		if label == "" {
			return field.Name
		}
		return label
	})
	myValidator.Validator = validatorObj
}

func (m myValidate) Validate(s interface{}) (err error) {
	err = m.Validator.Struct(s)
	if err == nil {
		return
	}
	err = m.DealErrMsg(err)
	return
}

func (m myValidate) DealErrMsg(err error) (errMsg error) {
	var (
		isExist bool
		field   string
	)
	errMsg = err

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return
	}

	for _, val := range err.(validator.ValidationErrors) {
		field = val.StructField()
		arr := strings.Split(val.Namespace(), ".")

		if len(arr) <= 0 {
			break
		}
		arr = strings.Split(arr[1], "|")
		if len(arr) <= 0 {
			break
		}
		if len(arr) == 1 {
			isExist = true
			errMsg = errors.New(arr[0])
			continue
		}
		for _, v := range arr {
			errMsgArr := strings.Split(v, ":")
			if errMsgArr[0] == val.Tag() && errMsgArr[1] != "" {
				errMsg = errors.New(errMsgArr[1])
				isExist = true
				break
			}
		}
	}

	if !isExist {
		errMsg = errors.New(fmt.Sprintf("请求参数 %s 校验不通过", field))
	}
	return
}
