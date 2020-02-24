package assert

import (
	"fmt"

	"github.com/kcarretto/paragon/pkg/script"
	"go.starlark.net/starlark"
)

// NoError will check if the passed value is a starlark.NoneType, if not it will error out the
// script.
//
//go:generate go run ../gendoc.go -lib assert -func noError -param errVal@starlark.Value -retval err@Error -doc "NoError will check if the passed value is a starlark.NoneType, if not it will error out the script."
//
// @callable:	assert.noError
// @param:		errVal	@starlark.Value
// @retval:		err 	@Error
//
// @usage:		assert.noError(err)
func NoError(errVal starlark.Value) error {
	_, ok := errVal.(starlark.NoneType)
	if !ok {
		errString, ok := errVal.(starlark.String)
		if !ok {
			return fmt.Errorf("assertion failed: error is not None")
		}
		return fmt.Errorf("assertion failed: error is not None: %v", errString)
	}
	return nil
}

func noError(parser script.ArgParser) (script.Retval, error) {
	errVal, err := parser.GetParam(0)
	if err != nil {
		return nil, err
	}

	retErr := NoError(errVal)
	return nil, retErr
}