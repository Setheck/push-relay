package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestC(t *testing.T) {
	c := NewConfig("testconfig", "testdata")
	o := struct {
		Token   string
		UserKey string
	}{}
	if err := c.Load("pushover", &o); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "testToken", o.Token)
	assert.Equal(t, "testUserKey", o.UserKey)
}
