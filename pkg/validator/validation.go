package validator

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// InitTrans 初始化翻译语言
func InitTrans(locale string) (trans ut.Translator, err error) {
	// 将 gin 的 validator 引擎改成 go-playground/validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 使用哪一个tag字段翻译
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("label")
		})

		zhTrans := zh.New()
		enTrans := en.New()
		// 创建，第一是参数是回退（备用），第二、三...个 支持参数
		uni := ut.New(enTrans, zhTrans)
		// 获取使用的翻译器实例
		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return nil, fmt.Errorf("初始化验证器 [%s] 翻译错误", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
