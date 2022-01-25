package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var v *validator.Validate

func init() {
	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v = engine
	} else {
		log.Fatalln("validator error")
	}
}
