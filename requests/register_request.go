package requests

import (
	"beego-bbs/models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

func ValidateRegistrationForm(valid *validation.Validation,data *models.User) map[string] []string{

	valid.Required(data.Name, "Name").Message("用户名不能为空" )
	valid.MinSize(data.Name, 6,"Name").Message("用户名最短为6位" )
	valid.Required(data.Password, "Password").Message("密码不能为空" )
	valid.MinSize(data.Password, 6,"Password").Message("密码最短为6位" )
	valid.Required(data.Password_confirm, "Password_confirm").Message("确认密码不能为空" )
	valid.MinSize(data.Password_confirm, 6,"Password_confirm").Message("确认密码最短为6位" )
	errors := make(map[string][]string)
	b, err := valid.Valid(data)
	if err != nil {
		fmt.Println(err,"xxxx")
	}
	if !b {

		for _, err := range valid.Errors {

			errors[err.Key] = append(errors[err.Key], err.Message)
		}

	}
	return errors
}