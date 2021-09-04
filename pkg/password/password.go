package password

import (
	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		logs.Error(err)
	}
	return string(bytes)
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	if err != nil {
		logs.Error(err)
	}
	return err==nil
}
// IsHashed 判断字符串是否是哈希过的数据
func IsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}