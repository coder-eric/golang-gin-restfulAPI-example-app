package validator

import (
	"fmt"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

// ValidateMultiEmails 校验以“,”分隔的多个邮箱 eg. test1@163.com,test2@163.com
func ValidateMultiEmails(fl validator.FieldLevel) bool {
	return isValidMultiEmails(fl.Field().String())
}

// ValidateMultiEmailsRegisterTranslationsFunc .
func ValidateMultiEmailsRegisterTranslationsFunc(ut ut.Translator) (err error) {
	if err = ut.Add("isValidMultiEmails", "{0}邮箱不合法，多个邮箱请以逗号(半角)分隔", false); err != nil {
		return
	}
	return

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		fmt.Println("警告: 翻译字段错误: ", fe)
		return fe.(error).Error()
	}
	return t
}

// 校验以“,”分隔的多个邮箱 eg. test1@163.com,test2@163.com
func isValidMultiEmails(mailsStr string) bool {
	if len(mailsStr) == 0 || mailsStr == "" {
		return false
	}
	mails := strings.Split(mailsStr, ",")
	for _, mail := range mails {
		if !isValidEmail(mail) {
			return false
		}
	}
	return true
}

var isEmailRe = regexp.MustCompile(`^([a-z_A-Z.0-9-])+@([a-zA-Z0-9_-])+\.([a-zA-Z0-9_-])+`)

// 判断是否正确的电子邮件
func isValidEmail(email string) bool {
	return isEmailRe.MatchString(email)
}
