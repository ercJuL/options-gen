// Code generated by options-gen. DO NOT EDIT.
package main

import (
	fmt461e464ebed9 "fmt"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptParamsSetter func(o *Params)

func NewParams(
	hash string,
	options ...OptParamsSetter,
) Params {
	o := Params{}

	// Setting defaults from field tag (if present)

	o.hash = hash

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func (o *Params) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("hash", _validate_Params_hash(o)))
	return errs.AsError()
}

func _validate_Params_hash(o *Params) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.hash, "hexadecimal"); err != nil {
		return fmt461e464ebed9.Errorf("field `hash` did not pass the test: %w", err)
	}
	return nil
}
