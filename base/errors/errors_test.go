package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	assertion := assert.New(t)

	assertion.True(Error{}.IsEmptyError())
	assertion.False(BadRequest.IsEmptyError())

	assertion.Equal("Bad Request", BadRequest.Error())

	assertion.Equal("BadRequest: Bad Request", BadRequest.String())

	err := BadRequest
	msg := "Customed Msg"
	err.SetMsg(msg)
	assertion.Equal(msg, err.Message)
}
