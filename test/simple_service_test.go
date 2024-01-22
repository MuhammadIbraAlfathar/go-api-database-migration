package test

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceFailed(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
