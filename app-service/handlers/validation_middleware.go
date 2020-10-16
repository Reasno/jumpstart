package handlers

import (
	"context"
	"github.com/Reasno/jumpstart/app-service/svc"
	"github.com/go-kit/kit/endpoint"
)

type Validator interface {
	validate() error
}

func validationMiddleware() svc.LabeledMiddleware {
	return func(name string, in endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			if t, ok := req.(Validator); ok {
				err = t.validate()
				if err != nil {
					return nil, newValidationError(err)
				}
			}
			resp, err = in(ctx, req)
			return
		}
	}
}

type ValidationError struct {
	err error
}

func (ve ValidationError) StatusCode() int {
	return 400
}
func (ve ValidationError) Error() string {
	return ve.err.Error()
}

func newValidationError(err error) ValidationError {
	return ValidationError{err}
}

