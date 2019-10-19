package app

import (
	"log"
	"testing"
)

func TestC(t *testing.T) {
	c := NewConfig("/home/seth/Github/push-relay")
	o := struct {
		Token   string
		UserKey string
	}{}
	if err := c.Load("pushover", &o); err != nil {
		t.Fatal(err)
	}
	log.Println(o)
}
