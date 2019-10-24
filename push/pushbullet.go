package push

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/setheck/push-relay/util"
)

type PushBulletPush struct {
	Type             string `json:"type"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Url              string `json:"url"`
	FileName         string `json:"file_name"`
	FileType         string `json:"file_type"`
	FileUrl          string `json:"file_url"`
	SourceDeviceIden string `json:"source_device_iden"`
	DeviceIden       string `json:"device_iden"`
	ClientIden       string `json:"client_iden"`
	ChannelTag       string `json:"channel_tag"`
	Email            string `json:"email"`
	Guid             string `json:"guid"`
}

const (
	pushesApi = "https://api.pushbullet.com/v2/pushes"
)

type PushBullet struct {
	AccessToken string
}

func (p *PushBullet) Send(m interface{}) (*http.Response, error) {
	msg, ok := m.(PushBulletPush)
	if !ok {
		return nil, errors.New("invalid type")
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Access-Token": p.AccessToken,
	}
	return util.PostJson(pushesApi, headers, bytes.NewBuffer(b))
}
