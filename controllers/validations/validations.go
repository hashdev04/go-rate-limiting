package validations

import (
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func Register() {
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("bookabledate", bookableDate)
	}
}
