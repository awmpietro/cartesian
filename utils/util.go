package utils

import "github.com/go-playground/validator/v10"

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func Validation(model interface{}) (err error) {
	var validate *validator.Validate
	validate = validator.New()
	err = validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}
