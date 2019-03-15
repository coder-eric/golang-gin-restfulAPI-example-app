package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	local_zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	translations_zh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    *ut.Translator
}

// 初始化gin验证器
func init() {
	binding.Validator = new(defaultValidator)
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			errs := err.(validator.ValidationErrors)
			messages := make([]string, len(errs), len(errs))
			for i, e := range errs {
				messages[i] = e.Translate(*v.trans)
			}
			return errors.New(strings.Join(messages, ", "))
		}
	}
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		fmt.Println("lazyinit validator")
		// 国际化
		localZH := local_zh.New()
		uni = ut.New(localZH, localZH)
		trans, _ := uni.GetTranslator("zh")

		v.validate = validator.New()
		v.trans = &trans
		// validate.v9 tag 默认validate，兼容老代码
		v.validate.SetTagName("validate")
		// 汉化验证提示
		translations_zh.RegisterDefaultTranslations(v.validate, trans)
		// 自定义验证器 https://godoc.org/gopkg.in/go-playground/validator.v9
		v.validate.RegisterValidation("isValidMultiEmails", ValidateMultiEmails)
		v.validate.RegisterTranslation("isValidMultiEmails", trans, ValidateMultiEmailsRegisterTranslationsFunc, translateFunc)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
