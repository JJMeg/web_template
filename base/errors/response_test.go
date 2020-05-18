package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResposne(t *testing.T) {
	assertion := assert.New(t)
	respErr := NewErrorResponse("mockId",BadRequest)

	assertion.Equal("mockId",respErr.RequestID)
	assertion.Equal(BadRequest.Code,respErr.GetStatusCode())
	}
