package test

import (
	"github.com/MuhammadIbraAlfathar/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, f := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	f()
}
