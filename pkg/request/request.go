package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	ContextWrapperService interface {
		Bind(a any) error
	}

	ContextWrapper struct {
		Context   echo.Context
		Validator *validator.Validate
	}
)

func NewContextWrapper(ctx echo.Context) ContextWrapperService {
	return &ContextWrapper{
		Context:   ctx,
		Validator: validator.New(),
	}
}

func (c *ContextWrapper) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		log.Printf("Error: Bind data failed %s", err.Error())
		return err
	}

	if err := c.Validator.Struct(data); err != nil {
		log.Printf("Error: Validate data failed: %s", err.Error())
		return err
	}

	return nil

}
