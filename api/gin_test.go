package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	testRouter := Setup()
	assert.True(t, testRouter.RedirectTrailingSlash)
	assert.True(t, testRouter.ForwardedByClientIP)
	assert.True(t, testRouter.UnescapePathValues)
	assert.False(t, testRouter.RedirectFixedPath)
	assert.False(t, testRouter.HandleMethodNotAllowed)
	assert.False(t, testRouter.RemoveExtraSlash)
}
