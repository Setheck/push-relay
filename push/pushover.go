package push

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/setheck/push-relay/util"
)

type PusherSender interface {
	Send(m interface{}) (*http.Response, error)
}

type PushoverMessage struct {
	// Required
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`

	// Optional
	Attatchment http.File `json:"attatchment,omitempty"` // TODO:(smt) file uploads
	Device      string    `json:"device,omitempty"`
	Title       string    `json:"title,omitempty"`
	Url         string    `json:"url,omitempty"`
	UrlTitle    string    `json:"url_title,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Sound       string    `json:"sound,omitempty"`
	Timestamp   int64     `json:"timestamp,omitempty"`
}

const (
	messagesApi = "https://api.pushover.net/1/messages.json"
)

type Pushover struct {
	Token   string
	UserKey string
}

func (p *Pushover) Send(m interface{}) (*http.Response, error) {
	msg, ok := m.(PushoverMessage)
	if !ok {
		return nil, errors.New("invalid type")
	}
	msg.Token = p.Token
	msg.User = p.UserKey
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return util.PostJson(messagesApi, nil, bytes.NewBuffer(b))
}
