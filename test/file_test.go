package test

import (
	"github.com/stretchr/testify/assert"
	"golang-restapi-httprouter/simple"
	"testing"
)

func TestConenction(t *testing.T) {
	connection, cleanup := simple.IntiliazeConnection("vandy")
	assert.NotNil(t, connection)
	cleanup()
}