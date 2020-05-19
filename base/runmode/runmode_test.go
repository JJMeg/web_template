package runmode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunMode(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(Production, RunMode("production"))
	assertion.Equal(Development, RunMode("development"))
	assertion.Equal(Test, RunMode("test"))

	assertion.True(Production.IsValid())
	assertion.True(Development.IsValid())
	assertion.True(Test.IsValid())
	assertion.False(RunMode("invalid").IsValid())

	assertion.True(Production.IsProduction())
	assertion.True(Development.IsDevelopment())
	assertion.True(Test.IsTest())
}
