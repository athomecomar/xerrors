package xerrors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type v1 struct{}

func (v *v1) IsValid(errors *Errors) {
	errors.Add("v1", "there's an error with v1")
}

type v2 struct{}

func (v *v2) IsValid(errors *Errors) {
	errors.Add("v2", "there's an error with v2")
}

func TestErrorsKeys(t *testing.T) {
	r := require.New(t)
	errors := Validate(&v1{}, &v2{})
	r.Contains(errors.Keys(), "v1")
	r.Contains(errors.Keys(), "v2")
}
