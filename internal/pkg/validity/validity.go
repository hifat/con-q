package validity

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Trans ut.Translator

func Register() {
	binding.Validator.Engine().(*validator.Validate).RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		Trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, Trans)
	}
}

func Validate(err error) (*validator.ValidationErrorsTranslations, error) {
	if _, ok := err.(validator.ValidationErrors); !ok {
		return nil, err
	}

	objErr := make(validator.ValidationErrorsTranslations)
	for _, e := range err.(validator.ValidationErrors) {
		objErr[e.Field()] = e.Translate(Trans)
	}

	return &objErr, nil
}
